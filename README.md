# Monitoring Check

## Description
This project provides a function to send alerts to a Mattermost channel when there have been no requests for a specified duration. It uses the Mattermost API to send a message containing information about the duration of inactivity.

## Installation
To install the project, follow these steps:

1. Clone the repository: `git clone https://github.com/dogaakcinar/monitoring-check.git`
2. Navigate to the project directory: `cd monitoring-check`
3. Install dependencies: `go get -d ./...`

## Usage
To use the project, follow these steps:

1. Import the `alert` package: `import "github.com/dogaakcinar/monitoring-check/alert"`
2. Call the `SendMattermostAlert` function with the desired duration: `alert.SendMattermostAlert(5 * time.Minute)`

## Dependencies
This project has the following dependencies:

- Go (version X.X.X)
- Mattermost server (version X.X.X)

To install the Mattermost server, refer to the official Mattermost documentation.

## Contributing
Contributions to this project are welcome! To contribute, follow these steps:

1. Fork the repository.
2. Create a new branch: `git checkout -b feature/my-feature`.
3. Make your changes and commit them: `git commit -am "Add my feature"`.
4. Push to the branch: `git push origin feature/my-feature`.
5. Open a pull request.

All pull requests will be reviewed by the project maintainers. Please ensure that your code follows the project's coding standards and includes appropriate tests.

## License
This project is licensed under the [MIT License](https://opensource.org/licenses/MIT).

## Contact
For any questions or concerns regarding this project, please create an issue on the [GitHub repository](https://github.com/dogaakcinar/monitoring-check/issues).
