from __future__ import annotations

import http.server
import pathlib
import tempfile
import threading
import unittest

import check_tutorial_links


class Handler(http.server.BaseHTTPRequestHandler):
    def do_GET(self) -> None:
        if self.path == "/ok":
            self.send_response(200)
            self.end_headers()
            self.wfile.write(b"ok")
            return
        if self.path == "/redirect":
            self.send_response(302)
            self.send_header("Location", "/ok")
            self.end_headers()
            return
        self.send_response(404)
        self.end_headers()

    def log_message(self, format: str, *args: object) -> None:
        return


class CheckTutorialLinksTest(unittest.TestCase):
    @classmethod
    def setUpClass(cls) -> None:
        cls.server = http.server.ThreadingHTTPServer(("127.0.0.1", 0), Handler)
        cls.thread = threading.Thread(target=cls.server.serve_forever, daemon=True)
        cls.thread.start()
        cls.base_url = f"http://127.0.0.1:{cls.server.server_port}"

    @classmethod
    def tearDownClass(cls) -> None:
        cls.server.shutdown()
        cls.server.server_close()
        cls.thread.join()

    def test_extract_links_reports_source_lines_and_deduplicates(self) -> None:
        with tempfile.TemporaryDirectory() as directory:
            path = pathlib.Path(directory) / "tutorial.rst"
            path.write_text(
                f"`first <{self.base_url}/ok>`_\n"
                f"`again <{self.base_url}/ok>`_\n"
                f"`redirect <{self.base_url}/redirect>`_\n",
                encoding="utf-8",
            )

            links = check_tutorial_links.extract_links(path)

        self.assertEqual([1, 2], [item.line for item in links[f"{self.base_url}/ok"]])
        self.assertIn(f"{self.base_url}/redirect", links)

    def test_check_link_accepts_success_and_redirect(self) -> None:
        success = check_tutorial_links.check_link(f"{self.base_url}/ok", timeout=1, retries=0)
        redirect = check_tutorial_links.check_link(
            f"{self.base_url}/redirect", timeout=1, retries=0
        )

        self.assertTrue(success.ok)
        self.assertTrue(redirect.ok)
        self.assertEqual(f"{self.base_url}/ok", redirect.final_url)

    def test_check_link_rejects_missing_page(self) -> None:
        result = check_tutorial_links.check_link(
            f"{self.base_url}/missing", timeout=1, retries=0
        )

        self.assertFalse(result.ok)
        self.assertEqual(404, result.status)


if __name__ == "__main__":
    unittest.main()
