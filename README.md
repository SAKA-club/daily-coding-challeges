# Daily Coding Challenges
Fine tune your skills by completing daily coding challenges

## Upcoming Features
- [ ] Slack integration
- [ ] Workflow to check that solution is by correct person (github handle). No other changes are made like to the env file, template,  etc.

## Prerequisites
1. Set your GH username in the `.env` file

## Complete challenges
- Read the challenge for the day
    ```
    go run main.go challenge
    ```
- Init the file for your solution. Creates a file under `<month>/<day>/<github-username>.go` with the template for your function.
    ```
    go run main.go init
    ```
- Test your solution
    ```
    go run main.go validate
    ```