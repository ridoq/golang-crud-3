package service

import (
	"base-gin/config"
	"base-gin/repository"
)

var (
	accountService   *AccountService
	personService    *PersonService
	publisherService *PublisherService
)

func SetupServices(cfg *config.Config) {
	accountService = NewAccountService(cfg, repository.GetAccountRepo())
	personService = NewPersonService(repository.GetPersonRepo())
	publisherService = NewPublisherService(repository.GetPublisherRepo())
}

func GetAccountService() *AccountService {
	if accountService == nil {
		panic("account service is not initialised")
	}

	return accountService
}

func GetPersonService() *PersonService {
	return personService
}

func GetPublisherService() *PublisherService {
	return publisherService
}
