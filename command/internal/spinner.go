package internal

import (
	"github.com/briandowns/spinner"
	"io"
	"time"
)

type Spinner struct {
	s *spinner.Spinner
}

type CharType int

const (
	CharSetDownload CharType = iota
	CharSetUpload
	CharSetProgress
)

var charSets = map[CharType][]string{
	CharSetDownload: {
		"[                    ]",
		"[                  <=]",
		"[                 <==]",
		"[                <===]",
		"[               <====]",
		"[              <=====]",
		"[             <======]",
		"[            <=======]",
		"[           <========]",
		"[          <=========]",
		"[         <==========]",
		"[        <===========]",
		"[       <============]",
		"[      <=============]",
		"[     <==============]",
		"[    <===============]",
		"[   <================]",
		"[  <=================]",
		"[ <==================]",
		"[<===================]",
	},
	CharSetUpload: {
		"[                    ]",
		"[=>                  ]",
		"[==>                 ]",
		"[===>                ]",
		"[====>               ]",
		"[=====>              ]",
		"[======>             ]",
		"[=======>            ]",
		"[========>           ]",
		"[=========>          ]",
		"[==========>         ]",
		"[===========>        ]",
		"[============>       ]",
		"[=============>      ]",
		"[==============>     ]",
		"[===============>    ]",
		"[================>   ]",
		"[=================>  ]",
		"[==================> ]",
		"[===================>]",
	},
	CharSetProgress: spinner.CharSets[9],
}

func NewSpinner(prefix string, compMessage string, charType CharType, out io.Writer) *Spinner {
	s := spinner.New(charSets[charType], 100*time.Millisecond)

	s.Writer = out
	s.Prefix = prefix
	s.FinalMSG = compMessage
	return &Spinner{
		s: s,
	}
}

func (s *Spinner) Start() {
	s.s.Start()
}

func (s *Spinner) Stop() {
	s.s.Stop()
}
