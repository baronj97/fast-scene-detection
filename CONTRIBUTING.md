# Contributing

If you wish to contribute to this project, please follow the guidelines below.

## Programming Language and Tools

This project is implemented in the Go programming language and will use any flavor of Go 1.13.

This project also uses the GoCV package, which requires an install of openCV. Please ensure openCV4 is installed.

## Workflow

When working on an issue, please follow these steps.

1. Select an issue you want to work on. These can be found under the "Issues" tab on Github.

1. Assign yourself. You may also add labels as necessary.

1. Create a new branch off of master by doing the following:
    ```
    git checkout master
    git pull
    git checkout -b your-name/issue
    ```

1. Work on your branch and when you are ready for review, update the CHANGELOG file by adding a new entry of the format:
    `- [Issue #issue#](https://github.com/baronj97/fast-scene-detection/issues/issue#) - Issue title`

1. Finally, create a pull request through the github website UI and select a reviewer. Once it is approved, merge it into master and delete your branch.

