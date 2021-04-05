# Get Started

```
git clone https://github.com/Fakorede/go-rest-api.git .

cd go-rest-api

go mod download
```

# Setup environment variables

> Add a .env file to the root with the below config

```
DB_HOST=
DB_PORT=
DB_USER=
DB_PASS=
DB_DATABASE=

JWT_SECRET=
```


# For Live reload

```
alias air='$(go env GOPATH)/bin/air'

air
```

# Features/Functionalities
- Login, Register, Forgot/Reset Password
- httpOnly cookies to store JWT
- Public and Secure Routes
- Authorization using Roles and Permissions
- Middlewares
- Relationships (One-to-Many and Many-to-Many)
- Migrations
- Pagination
- Mail for Reset Password Link([MailHog](https://github.com/mailhog/MailHog) must be installed locally and running to catch mails)
- Image Upload
- CSV Export

# Routes

can be found in `routes/routes.go`

```
...
```

# Built with

- Go
- Go Fiber
- GORM
- MySql