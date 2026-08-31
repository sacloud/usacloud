#!/usr/bin/env python3

"""Check external links in API tutorial examples."""

from __future__ import annotations

import argparse
import concurrent.futures
import dataclasses
import pathlib
import re
import sys
import time
import urllib.error
import urllib.request


URL_PATTERN = re.compile(r"https?://[^\s<>()\[\]{}\"']+")
TRAILING_PUNCTUATION = ".,;:!?`"
DEFAULT_ASSET_PATTERN = "tutorial-example-*-api.rst"
USER_AGENT = "usacloud-tutorial-link-checker/1.0"


@dataclasses.dataclass(frozen=True)
class Occurrence:
    path: pathlib.Path
    line: int


@dataclasses.dataclass(frozen=True)
class LinkResult:
    url: str
    status: int | None
    final_url: str | None
    error: str | None

    @property
    def ok(self) -> bool:
        return self.error is None and self.status is not None and 200 <= self.status < 400


def extract_links(path: pathlib.Path) -> dict[str, list[Occurrence]]:
    links: dict[str, list[Occurrence]] = {}
    for line_number, line in enumerate(path.read_text(encoding="utf-8").splitlines(), start=1):
        for match in URL_PATTERN.finditer(line):
            url = match.group(0).rstrip(TRAILING_PUNCTUATION)
            links.setdefault(url, []).append(Occurrence(path=path, line=line_number))
    return links


def check_link(url: str, timeout: float, retries: int) -> LinkResult:
    request = urllib.request.Request(
        url,
        headers={
            "User-Agent": USER_AGENT,
            "Accept": "text/html,application/xhtml+xml,*/*;q=0.8",
        },
        method="GET",
    )

    for attempt in range(retries + 1):
        try:
            with urllib.request.urlopen(request, timeout=timeout) as response:
                response.read(1)
                return LinkResult(
                    url=url,
                    status=response.status,
                    final_url=response.geturl(),
                    error=None,
                )
        except urllib.error.HTTPError as error:
            status = error.code
            final_url = error.geturl()
            reason = error.reason
            error.close()
            if status < 500 or attempt == retries:
                return LinkResult(
                    url=url,
                    status=status,
                    final_url=final_url,
                    error=f"HTTP {status} {reason}",
                )
        except (urllib.error.URLError, TimeoutError) as error:
            if attempt == retries:
                return LinkResult(
                    url=url,
                    status=None,
                    final_url=None,
                    error=str(error.reason if isinstance(error, urllib.error.URLError) else error),
                )
        time.sleep(0.5 * (attempt + 1))

    raise AssertionError("unreachable")


def default_paths() -> list[pathlib.Path]:
    assets = pathlib.Path(__file__).resolve().parent.parent / "assets"
    return sorted(assets.glob(DEFAULT_ASSET_PATTERN))


def parse_args(argv: list[str]) -> argparse.Namespace:
    parser = argparse.ArgumentParser(
        description="Check HTTP(S) links in usacloud API tutorial examples.",
    )
    parser.add_argument(
        "paths",
        nargs="*",
        type=pathlib.Path,
        help=f"RST files to check (default: assets/{DEFAULT_ASSET_PATTERN})",
    )
    parser.add_argument("--timeout", type=float, default=15.0, help="Timeout per request in seconds")
    parser.add_argument("--retries", type=int, default=1, help="Retries for server and network errors")
    parser.add_argument("--jobs", type=int, default=4, help="Maximum concurrent requests")
    args = parser.parse_args(argv)
    if args.timeout <= 0:
        parser.error("--timeout must be greater than zero")
    if args.retries < 0:
        parser.error("--retries must not be negative")
    if args.jobs <= 0:
        parser.error("--jobs must be greater than zero")
    return args


def main(argv: list[str] | None = None) -> int:
    args = parse_args(argv if argv is not None else sys.argv[1:])
    paths = args.paths or default_paths()
    missing_paths = [path for path in paths if not path.is_file()]
    if missing_paths:
        for path in missing_paths:
            print(f"ERROR: file not found: {path}", file=sys.stderr)
        return 2
    if not paths:
        print(f"ERROR: no files matched {DEFAULT_ASSET_PATTERN}", file=sys.stderr)
        return 2

    links: dict[str, list[Occurrence]] = {}
    for path in paths:
        for url, occurrences in extract_links(path).items():
            links.setdefault(url, []).extend(occurrences)
    if not links:
        print("ERROR: no HTTP(S) links found", file=sys.stderr)
        return 2

    with concurrent.futures.ThreadPoolExecutor(max_workers=args.jobs) as executor:
        futures = {
            executor.submit(check_link, url, args.timeout, args.retries): url for url in links
        }
        results = [future.result() for future in concurrent.futures.as_completed(futures)]

    failures = 0
    for result in sorted(results, key=lambda item: item.url):
        if result.ok:
            redirect = ""
            if result.final_url and result.final_url != result.url:
                redirect = f" -> {result.final_url}"
            print(f"OK {result.status}: {result.url}{redirect}")
            continue

        failures += 1
        print(f"BROKEN: {result.url}: {result.error}", file=sys.stderr)
        for occurrence in links[result.url]:
            print(f"  {occurrence.path}:{occurrence.line}", file=sys.stderr)

    print(f"Checked {len(links)} unique links in {len(paths)} files; {failures} broken.")
    return 1 if failures else 0


if __name__ == "__main__":
    raise SystemExit(main())
