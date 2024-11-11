package repository

import "base-gin/storage"

var (
	accountRepo   *AccountRepository
	personRepo    *PersonRepository
	publisherRepo *PublisherRepository
	bookRepo      *BookRepository
	authorRepo    *AuthorRepository
	borrowingRepo    *BorrowingRepository
)

func SetupRepositories() {
	db := storage.GetDB()
	accountRepo = NewAccountRepository(db)
	personRepo = NewPersonRepository(db)
	publisherRepo = NewPublisherRepository(db)
	bookRepo = NewBookRepository(db)
	authorRepo = NewAuthorRepository(db)
	borrowingRepo = NewBorrowingRepository(db)
}

func GetAccountRepo() *AccountRepository {
	return accountRepo
}

func GetPersonRepo() *PersonRepository {
	return personRepo
}

func GetPublisherRepo() *PublisherRepository {
	return publisherRepo
}

func GetBookRepo() *BookRepository {
	return bookRepo
}

func GetAuthorRepo() *AuthorRepository {
	return authorRepo
}
func GetBorrowingRepo() *BorrowingRepository {
	return borrowingRepo
}
