// Copyright (c) seasonjs. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package sd

import (
	"fmt"
	"os"
)

func dumpSDLibrary(gpu bool) (*os.File, error) {
	file, err := os.CreateTemp("", libName)
	if err != nil {
		return nil, fmt.Errorf("error creating temp file: %w", err)
	}

	if err := os.WriteFile(file.Name(), getDl(gpu), 0400); err != nil {
		return nil, fmt.Errorf("error writing file: %w", err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			return
		}
	}()
	return file, nil
}
