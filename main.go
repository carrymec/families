package main

import (
	"context"
	"encoding/json"
	"fmt"

	"log"
	"os"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"

	"github/carrymec/families/configs"
	"github/carrymec/families/logger"
	"github/carrymec/families/person"
	"github/carrymec/families/relationship"
)

func main() {
	r := gin.New()
	readFile, err := os.ReadFile("config.json")
	if err != nil {
		fmt.Printf("read config.json err: %#v\n", err)
		panic(err)
	}
	var cfg configs.Config
	err = json.Unmarshal(readFile, &cfg)
	if err != nil {
		panic(err)
	}
	if err := logger.InitLogger(cfg.LogConfig); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		panic(err)
	}

	gin.SetMode(configs.Conf.Mode)

	ctx := context.Background()
	dbUri := fmt.Sprintf("neo4j://%s", cfg.Neo4jConfig.Url)
	driver, err := neo4j.NewDriverWithContext(
		dbUri,
		neo4j.BasicAuth(cfg.Neo4jConfig.User, cfg.Neo4jConfig.Password, ""))
	defer func(driver neo4j.DriverWithContext, ctx context.Context) {
		err := driver.Close(ctx)
		if err != nil {
			panic(err)
		}
	}(driver, ctx)

	err = driver.VerifyConnectivity(ctx)
	if err != nil {
		panic(err)
	}
	logger.Logger.Info("connection success...")
	session := driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: cfg.Neo4jConfig.DatabaseName})
	defer func(session neo4j.SessionWithContext, ctx context.Context) {
		err := session.Close(ctx)
		if err != nil {
			panic(err)
		}
	}(session, ctx)
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	dao := person.NewPersonDao(logger.Logger, session)
	service := person.NewService(logger.Logger, dao)
	controller := person.NewPersonController(logger.Logger, service)
	controller.Register(r)

	relationDao := relationship.NewRelationDao(logger.Logger, session)
	relationService := relationship.NewService(logger.Logger, relationDao)
	relationController := relationship.NewRelationController(logger.Logger, relationService)
	relationController.Register(r)

	if err := endless.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), r); err != nil {
		log.Fatalf("listen: %s\n", err)
	}

	fmt.Println("Server exiting")
}
