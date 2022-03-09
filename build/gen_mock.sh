#need to install mockery
go install github.com/vektra/mockery/v2@latest
# or
# brew install mockery

mockery --recursive --name=CompanyStore
mockery --recursive --name=EmployeeStore
mockery --recursive --name=ProjectStore