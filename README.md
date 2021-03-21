# Deal Maker (Better market)
This project adapted a hopefully novel method of developing a REST-API. The method is currently named "Factory - Streamline - Conveyor belt" pattern. Document about the pattern can be found in https://github.com/itzMeerkat/streamline. Please read the `README` of that repo first before continue.

## Project Status
Current version of dealmaker should work with v0.0.1 of both `streamline` and `mentally-friendly-infra`. Please change `go.mod` yourself and have a try!

## Folders
+ dal: Database connection initialization and some related constants.
+ factory: To assemble streamlines. In the very long future this will be replaced by code generation.
+ handler: Each file should contain exactly one function, matching the `gin` handler function signature. Functions have a one-one relation to APIs.
+ model: Your business data model definitions.
+ procedure: Contains multiple sub-modules. Each one is an individual procedure. As for to what extend it should be considered a procedure? I don't have a conclusion yet, please discuss your thought with me.
+ resp_def: Define all `JSON` responses. (Not completed yet)
+ shared: Some procedures I think could be used in other projects, will become standalone package when they reach a stable state.

## Naming
+ No plural. It will introduce tons of inconsistency in names, and programmers can easily identify whether it is plural or singular by data type.

## Contributing
Do NOT push to master directly. Each person should checkout a personal working branch. After push modifications to YOUR branch, create a merge request that merge into `master` branch.

Do NOT push any unnecessary file to the repo. Including `.DS_Store`, `.vscode`, `.idea` etc..

After the Git assignment I have faith in you.

## Note
Currently, `procedure`s in `shared` folder is more complete. If you need reference, please go there instead of `procedure` folder.

## Note2
To get this running, you need to modify `go.mod` file, remove those `replace`s or change the path to your local path. I use replace is because those dependencies are also under active developing.

They are https://github.com/itzMeerkat/streamline and https://github.com/itzMeerkat/mentally-friendlly-infra.
