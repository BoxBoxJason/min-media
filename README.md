# min-media

FFmpeg optimizer script, reduces the size of videos, photos and audio files without losing quality.

This script searches recursively for video, photo and audio files in the current directory and its subdirectories.
It uses ffmpeg to optimize the files and reduce their size.
The optimized files are saved in the same directory and replace the original files if the optimization is successful.
The script also creates a log file to keep track of the optimization process and any errors that occur.

## Usage

### Main Command

The main command is `min-media`, which can be used with various flags and subcommands.
It can run the entire optimization process or specific subcommands for more granular control.


```plaintext
Usage:
  min-media [flags]
  min-media [command]

Aliases:
  min-media, optimize

Available Commands:
  clear       Delete the optimized files without saving them
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  overwrite   Overwrite the original files with the optimized files
  status      Report the status of the optimization process

Flags:
  -a, --all       Optimize all files
  -A, --audio     Optimize audio files
  -d, --dry-run   Perform a dry run without modifying any files
  -f, --force     Overwrite the original files with the optimized files
  -h, --help      help for min-media
  -I, --image     Optimize image files
  -q, --quiet     Enable quiet output
  -v, --verbose   Enable verbose output
      --version   version for min-media
  -V, --video     Optimize video files

Use "min-media [command] --help" for more information about a command.
```

### Status

The `min-media status` command provides a summary of the optimization process.
It shows the files that were successfully optimized, those that failed, and those that were not processed.

```plaintext
Report the status of the optimization process
This command will show all files that were successfully optimized, those that failed and those that were not processed.

Usage:
  min-media status [flags]

Flags:
  -h, --help      help for status
      --version   version for status

Global Flags:
  -a, --all       Optimize all files
  -A, --audio     Optimize audio files
  -d, --dry-run   Perform a dry run without modifying any files
  -I, --image     Optimize image files
  -q, --quiet     Enable quiet output
  -v, --verbose   Enable verbose output
  -V, --video     Optimize video files
```

### Overwrite

The `min-media overwrite` command allows you to overwrite the original files with the optimized files.
This command will delete all files that were optimized by the script (those which end with `_optimized.ext`).

```plaintext
Overwrite the original files with the optimized files
Delete all files that were optimized by the script (those which end with _optimized.ext)

Usage:
  min-media overwrite [flags]

Flags:
  -h, --help      help for overwrite
      --version   version for overwrite

Global Flags:
  -a, --all       Optimize all files
  -A, --audio     Optimize audio files
  -d, --dry-run   Perform a dry run without modifying any files
  -I, --image     Optimize image files
  -q, --quiet     Enable quiet output
  -v, --verbose   Enable verbose output
  -V, --video     Optimize video files
```

### Clear

The `min-media clear` command allows you to delete the optimized files without saving them.
This command will delete all files that were optimized by the script (those which end with `_optimized.ext`) and will clear the status of the optimization process.

```plaintext
Delete all files that were optimized by the script (those which end with _optimized.ext), also resets the status of the optimization process

Usage:
  min-media clear [flags]

Flags:
  -h, --help      help for clear
      --version   version for clear

Global Flags:
  -a, --all       Optimize all files
  -A, --audio     Optimize audio files
  -d, --dry-run   Perform a dry run without modifying any files
  -I, --image     Optimize image files
  -q, --quiet     Enable quiet output
  -v, --verbose   Enable verbose output
  -V, --video     Optimize video files
```
