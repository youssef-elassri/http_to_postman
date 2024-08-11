#!/bin/bash

BINARY_NAME="http_to_postman"

go build -o $BINARY_NAME .

INSTALL_DIR="$HOME/.local/bin"

mkdir -p "$INSTALL_DIR"

mv $BINARY_NAME "$INSTALL_DIR"/

if ! grep -q "$INSTALL_DIR" ~/.zshrc && ! grep -q "$INSTALL_DIR" ~/.bash_profile; then
    if [ -n "$ZSH_VERSION" ]; then
        echo "export PATH=\"\$PATH:$INSTALL_DIR\"" >> ~/.zshrc
        source ~/.zshrc
    else
        echo "export PATH=\"\$PATH:$INSTALL_DIR\"" >> ~/.bash_profile
        source ~/.bash_profile
    fi
fi

echo "Installation complete! You can now use $BINARY_NAME directly."
