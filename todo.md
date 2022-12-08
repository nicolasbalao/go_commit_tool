# Note 

## General behavior 

1. List of type
2. Input scope (optional)
3. Input description
4. Text area body
5. Input ask breaking change
    If have breaking change write in the footer BREAKING CHANGE and add ! to the type like feat!
6. Input footer
7. Confirm commit and do the commit

## Structure of the app

- main => start the application
- tui/ => store all ui component 
    - tui.go => manage all ui component 

## Helper

- type: fix, feat, docs, build, chore, styke, refactor, perf, test

### Structure of the commit message

type string
description string
body string
breaking bool
footer string

## Todo

---2022-12-08---
- [x] set up the tui.go
- [x] create the type component
- [x] with the main
- [x] Make breaking component

## To improove

- [ ] breaking change with a confirm like gum



