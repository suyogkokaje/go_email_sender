# go_email_sender
**go_email_sender** will help you to send the emails in bulk

## Usage
1. Add the `.csv` file in the `data` directory and edit the path in the `csvPath` variable in the `main.go` file.
   Note that the `.csv` file must contain the column with name `email` (case sensitive).
  
2. Add the HTML template in the `template` directory and edit the path in the `templatePath` variable in the `main.go` file.
   Add the variables in the template file according to the instructions mentioned in the existing `template.html`.

3. Create `.env` file in the root directory of this project. Populate it with the variables from the `.env.example` file and replace them with your own credentials.
   
4. Specify the email subject in the `email_content.json` file.
   
5. Execute the command `go run main.go`

## Feature Requests and Contributions
- If you have any feature requests, please open an issue [here](https://github.com/suyogkokaje/go_email_sender/issues)
- Feel free to contribute to this project by raising a PR!
