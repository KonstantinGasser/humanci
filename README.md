# humanci 👉 🤖 👈

## First things first

`humanci` is most likely not a CLI-Framework to realistically replace any of the current ones. Why? Well because in `humanci` we don't know anything about `flags`, `args` nor `sub-commands`.
Welll kind of. With `humanci` you can build `CLI-Tools` which commands can be spelled in a fluent language. Let me give you an example. I sometimes need to check which web-server I forgot
to stop and is now block port 8080. However, I takes me 3 google searches and 4 Safari tabs to find the `lsof -i tcp:8080` command to release my pain. `humanci` should allow you to build something where you can just type in `which process is listening on 8080`. 

This whole things is meant to be fun project rather than a project to build the most efficient and performant CLI-Framework. If you are looking for one of those go check out [cobra](https://github.com/spf13/cobra), pretty awesome!