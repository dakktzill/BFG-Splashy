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

package db

import (
	"encoding/json"
	"errors"
	"io/fs"
	"os"
)

type Constants struct {
	StartingCoins    int64 `json:"starting_coins"`
	StartingDiamonds int64 `json:"starting_diamonds"`
	StartingFood     int64 `json:"starting_food"`
	StartingXP       int64 `json:"starting_xp"`
	StartingShards   int64 `json:"starting_shards"`
	StartingLevel    int   `json:"starting_level"`
}

func DefaultConstants() Constants {
	return Constants{
		StartingCoins:    750_000_000,
		StartingDiamonds: 1_000_000,
		StartingFood:     250_000_000,
		StartingXP:       0,
		StartingShards:   100_000_000,
		StartingLevel:    1,
	}
}

func LoadConstants(path string) (Constants, error) {
	raw, err := os.ReadFile(path)
	if errors.Is(err, fs.ErrNotExist) {
		return DefaultConstants(), nil
	}
	if err != nil {
		return DefaultConstants(), err
	}
	var c Constants
	if err := json.Unmarshal(raw, &c); err != nil {
		return DefaultConstants(), err
	}
	return c, nil
}
