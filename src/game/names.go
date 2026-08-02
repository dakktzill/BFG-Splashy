/* Project Encore: BFG - Localized Private Game Restoration Server
 * Copyright (C) 2026 Paficent <paficent@tutamail.com> & Contributors
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package game

import (
	"math/rand"
	"regexp"
	"strings"
	"unicode"
)

func randomMonsterName(names []string) string {
	if len(names) == 0 {
		return "Monster"
	}
	return names[rand.Intn(len(names))]
}

const nameAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!¨\"#$&'()*+,-./:;<=>?@}{0123456789|£©¿®`~^ÀÁÂÄÇÈÉÊËÌÍÎÏÑÒÓÔÖÙÚÛÜßàáâäçèéêëìíîïñòóôöùúûü_ÆæÃãÕõАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯабвгдеёжзийклмнопрстуфхцчшщъыьэюя€₽¡"

var nameAlphabetSet = func() map[rune]bool {
	m := make(map[rune]bool)
	for _, r := range nameAlphabet {
		m[r] = true
	}
	return m
}()

func sanitizeName(name string) string {
	if name == "" {
		return ""
	}
	var b strings.Builder
	for _, c := range name {
		switch {
		case nameAlphabetSet[c]:
			b.WriteRune(c)
		case unicode.IsSpace(c):
			b.WriteRune(' ')
		default:
			b.WriteRune('?')
		}
	}
	return b.String()
}

var leetReplacer = strings.NewReplacer(
	"@", "a", "4", "a", "3", "e", "1", "i", "!", "i",
	"0", "o", "$", "s", "5", "s", "7", "t",
)

var separatorRegex = regexp.MustCompile(`[\s._\-]+`)

func normalizeText(text string) string {
	return leetReplacer.Replace(strings.ToLower(text))
}

func collapseSeparators(text string) string {
	return separatorRegex.ReplaceAllString(text, "")
}

func collapseRepeats(text string) string {
	var b strings.Builder
	var prev rune
	first := true
	for _, r := range text {
		if first || r != prev {
			b.WriteRune(r)
		}
		prev = r
		first = false
	}
	return b.String()
}

// I wonder what Dave was thinking when he typed all of these out
var badWords = []string{
	"fuck", "fucker", "fucking", "motherfucker", "mf", "shit", "bullshit", "bitch", "bitches",
	"ass", "asshole", "dick", "dildo", "cock", "cocksucker", "pussy", "pussies", "slut",
	"whore", "cum", "cumming", "jizz", "jerkoff", "handjob", "blowjob", "boob", "boobs",
	"tits", "tit", "nipple", "porn", "porno", "pornhub", "sex", "sexy", "s3x", "suck",
	"sucking", "deepthroat", "anal", "anus", "buttsex", "butthole", "balls", "testicles",
	"scrotum", "masturbate", "masturbation", "orgasm", "orgy", "fetish", "bdsm",
	"bondage", "spank", "spanking", "horny", "hentai", "rule34",
	"idiot", "moron", "dumbass", "retard", "retarded", "stupid", "loser", "noob",
	"trash", "garbage", "clown", "scumbag", "dipshit", "douche", "douchebag",
	"jackass", "prick", "tool", "twat", "wanker", "shithead", "dirtbag",
	"kill", "kys", "die", "suicide", "murder", "rapist", "rape", "raping",
	"terrorist", "bomb", "massacre", "genocide", // god forbid children are exposed to the evil of humanity early
	"cocaine", "heroin", "meth", "weed", "marijuana", "crack", "lsd", "drugdealer", // I see salvia is exempt
	"fag", "faggot",
	"fuk", "fuc", "phuck", "fucc", "fuq", "shiit", "sh1t", "b1tch", "biatch",
	"azzhole", "a55hole", "d1ck", "d!ck", "c0ck", "p0rn", "s3xy", "f@g", "f@ggot",
	"satan", "devil", "racist", "bigot",
} // haha!

func containsBadWord(text string) bool {
	if text == "" {
		return false
	}
	text = normalizeText(text)
	collapsed := collapseRepeats(collapseSeparators(text))
	for _, word := range badWords {
		if strings.Contains(text, word) || strings.Contains(collapsed, word) {
			return true
		}
	}
	return false
}

var whitespaceOnly = regexp.MustCompile(`^\s*$`)

func invalidName(name string) string {
	if name == "" {
		return "INVALID_DISPLAY_NAME"
	}
	if strings.Contains(name, "%") {
		return "INVALID_CHAR_DISPLAY_NAME"
	}
	if strings.Contains(name, "<c") {
		return "INVALID_CHAR_DISPLAY_NAME"
	}
	if strings.Contains(name, "</") {
		return "INVALID_CHAR_DISPLAY_NAME"
	}
	if containsBadWord(name) {
		return "BAD_WORD_DISPLAY_NAME"
	}
	if whitespaceOnly.MatchString(name) {
		return "INVALID_WHITESPACE_DISPLAY_NAME"
	}
	return ""
}
