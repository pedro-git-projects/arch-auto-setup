run:
	go build . && ./arch-auto-setup

clean:
	rm -rf ./arch-dotfiles/ ./yay ./nvim-dotfiles/
