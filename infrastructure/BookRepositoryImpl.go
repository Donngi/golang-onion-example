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
	Id        string `dynamodbav:"id"`
	Name      string `dynamodbav:"name"`
	Author    string `dynamodbav:"author"`
	CreatedAt string `dynamodbav:"createdAt"`
	UpdatedAt string `dynamodbav:"updatedAt"`
}

func NewBookRepositoryImpl(client *dynamodb.Client) *BookRepositoryImpl {
	return &BookRepositoryImpl{client: client}
}

func (repo *BookRepositoryImpl) Add(book *domain.Book) (*domain.Book, error) {
	now := time.Now().Format(time.RFC3339)
	item, err := attributevalue.MarshalMap(&BookTableItem{
		Id:        book.Id,
		Name:      book.Name,
		Author:    book.Author,
		CreatedAt: now,
		UpdatedAt: now,
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

	resItem := &BookTableItem{}
	err = attributevalue.UnmarshalMap(res.Attributes, resItem)
	if err != nil {
		return nil, err
	}

	return &domain.Book{
		Id:     resItem.Id,
		Name:   resItem.Name,
		Author: resItem.Author,
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
		Id:     resItem.Id,
		Name:   resItem.Name,
		Author: resItem.Author,
	}, nil
}
