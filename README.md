## Arch Auto Setup

This script automates the setup of system preferences after a clean Arch Linux installation.

### Dependencies

- git
- golang

### Manual Steps

Before running the script, you need to perform the following manual steps:

#### Pre-install

1. Add user to the wheel group:
    ```bash
    usermod -aG wheel <username>
    ```
    Then use `visudo` and uncomment the line `%wheel ALL=(ALL) ALL`.

2. Generate default folders:
    Install `xdg-user-dirs` and run:
    ```bash
    xdg-user-dirs-update
    ```

#### Post-install

1. Set the theme to Gruvbox-Material-Dark and the icon theme to oomox-gruvbox-dark using lxappearance.

2. Compile Spotify-tray and Vesktop from source, then create symbolic links:
    ```bash
    ln -s $HOME/path/to/compiled /usr/local/bin/name-for-dmenu
    ```

### Usage

Clone this repository and run the script:

```bash
git clone https://github.com/pedro-git-projects/arch-auto-setup.git
cd arch-auto-setup 
go build .
./arch-auto-setup pkg
```

### Note

Ensure that you run the script with appropriate permissions, especially when modifying system directories.
