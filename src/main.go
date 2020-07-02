// Copyright 2010 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"

	"github.com/ahmdrz/goinsta/v2"
)

func main() {
	insta := goinsta.New("tiburon.new", "skiminok")
	if err := insta.Login(); err != nil {
		switch v := err.(type) {
		case goinsta.ChallengeError:
			err := insta.Challenge.Process(v.Challenge.APIPath)
			if err != nil {
				log.Fatalln(err)
			}

			err = insta.Challenge.SendSecurityCode("546 309")
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
	defer insta.Logout()
	fm, err := insta.Timeline.Stories()
	if err != nil {
		log.Fatalln(err)
	}
	st := fm.Stories
	fmt.Printf("%+v", st)
}
