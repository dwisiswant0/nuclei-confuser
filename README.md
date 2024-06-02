# nuclei confuser

This repository gathers matchers from Nuclei templates designed to fool the Nuclei scanner.

## Manual

> [!NOTE]  
> The following tools are required:
> * [go](https://go.dev/doc/install).
> * [task](https://github.com/go-task/task) runner.
> * [yq](https://github.com/mikefarah/yq) (YAML processor).

* Open the `.env` file and modify the `NUCLEI_DIR` variable to point to the absolute path of your local Nuclei templates repository. _(Optional)_
* Run `task dependencies` to install all necessary dependencies required.
* Run `task dump-all-matchers` to extract and compile both words and regex matchers from the Nuclei templates.
* Run `task clean-build` to clean the build directory.
* Run `task build-all` to build all matchers.

## Build Structure

You can find the build tree in the [`build/`](/build) directory.

Each file is prefixed as follows:

* `words`: contains matchers of the words type.
* `regex`: contains matchers of the regex type.
* `combined`: contains a combination of matchers from both words and regex types.

## Usage

Simply embed the desired file into your footer or wherever you like.

## License

Released under `DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE`.