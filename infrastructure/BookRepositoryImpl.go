package infrastructure

import (
	"context"
	"time"

	"github.com/Donngi/golang-onion-example/domain"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type BookRepositoryImpl struct {
	client *dynamodb.Client
}

type BookTableItem struct {
	id        string `dynamodbav:"id"`
	name      string `dynamodbav:"name"`
	author    string `dynamodbav:"author"`
	createdAt string `dynamodbav:"createdAt"`
	updatedAt string `dynamodbav:"updatedAt"`
}

func NewBookRepositoryImpl(client *dynamodb.Client) *BookRepositoryImpl {
	return &BookRepositoryImpl{client: client}
}

func (repo *BookRepositoryImpl) Add(book *domain.Book) (*domain.Book, error) {
	now := time.Now().String()
	item, err := attributevalue.MarshalMap(&BookTableItem{
		id:        book.Id,
		name:      book.Name,
		author:    book.Author,
		createdAt: now,
		updatedAt: now,
	})
	if err != nil {
		return nil, err
	}

	res, err := repo.client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("Books"),
		Item:      item,
	})
	if err != nil {
		return nil, err
	}

	resItem := BookTableItem{}
	err = attributevalue.UnmarshalMap(res.Attributes, resItem)
	if err != nil {
		return nil, err
	}

	return &domain.Book{
		Id:     resItem.id,
		Name:   resItem.name,
		Author: resItem.author,
	}, nil
}

func (repo *BookRepositoryImpl) Get(id string) (*domain.Book, error) {
	res, err := repo.client.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String("Books"),
		Key: map[string]types.AttributeValue{
			id: &types.AttributeValueMemberS{
				Value: id,
			},
		}})
	if err != nil {
		return nil, err
	}

	resItem := BookTableItem{}
	err = attributevalue.UnmarshalMap(res.Item, resItem)
	if err != nil {
		return nil, err
	}

	return &domain.Book{
		Id:     resItem.id,
		Name:   resItem.name,
		Author: resItem.author,
	}, nil
}
