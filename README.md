# Harmony Server
## A Free And Open Source Server For A Communications Platform With Guilds

![Harmony Logo](https://avatars3.githubusercontent.com/u/58618825?s=200)

## Setup

### For setup you have two options : 

1. Use `docker-compose` to have the server automatically configured for you
    
    This is the simplest option. Simply `git clone` this repository, install `docker` and `docker-compose`, then run `docker-compose up` and let docker do all of the legwork for you!

    **important note : you might have issues with github or alpine repositories refusing connections, just rerun the command until it works. (Looking for a fix)**
2. Setup Server Manually

    This is tougher but more desirable for an interactive development environment
    
    It involves you installing **postgresql** on your machine, modifying `.env` to match your postgres configuration, and adding the `harmony` table.

    Then, you must install Golang. Once that is done, it should only be a matter of running `go build && ./harmony-server` on Linux/Mac and `go build && harmony-server.exe` on Windows. 

Expected Result
![Server Preview](https://i.imgur.com/NaDOxYX.png)

**Recommended Tools And Configurations**

 - **DBeaver** to view the postgresql database
 - **VSCode** for editing Golang. (Use the Go extension)
    
    Note : Use this to not get garbage experience : 
    ```json
    "go.useLanguageServer": true,
    "[go]": {
        "editor.snippetSuggestions": "none",
        "editor.formatOnSave": true,
        "editor.codeActionsOnSave": {
            "source.organizeImports": true,
        }
    },
    "gopls": {
        "usePlaceholders": true, // add parameter placeholders when completing a function
        // Experimental settings
        "completeUnimported": true, // autocomplete unimported packages
        "deepCompletion": true,     // enable deep completion
    },
    "go.languageServerExperimentalFeatures": {
        "format": true,
        "autoComplete": true,
        "rename": true,
        "goToDefinition": true,
        "hover": true,
        "signatureHelp": true,
        "goToTypeDefinition": true,
        "goToImplementation": true,
        "documentSymbols": true,
        "workspaceSymbols": true,
        "findReferences": true,
        "diagnostics": true,
        "documentLink": true
    },
    ```