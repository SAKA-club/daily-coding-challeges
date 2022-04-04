# Daily Coding Challenges
Fine tune your skills by completing daily coding challenges

## Upcoming Features
- [ ] Slack integration
- [ ] Workflow to check that solution is by correct person (github handle). No other changes are made like to the env file, template,  etc.

## Prerequisites
- Set your GH username in the `.env` file
   ```shell
   USERNAME=<GITHUB USERNAME>
   ```

## Complete challenges
- Read the challenge for the day
    ```shell
    go run main.go list
    ```
- Init the file for your solution. Creates a file under `<month>/<day>/<github-username>.go` with the template for your function.
    ```shell
    go run main.go init
    ```
- Test your solution
    ```shell
    go run main.go test
    ```