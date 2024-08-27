package main

import (
  "context"
  "fmt"
  "os"
  gpt3 "github.com/PullRequestInc/go-gpt3"
)

type openAI struct {
  Ctx context.Context
  Cli gpt3.Client
}

func NewAI() *openAI {
  return &openAI{}
}

func (ai *openAI) init() {
  ai.Ctx = context.Background()
  apiKey := os.Getenv("APIKEY")
  if apiKey == "" {
    panic("Set APIKEY=API-Key")
  }
  ai.Cli = gpt3.NewClient(apiKey)
}

func (ai openAI) printResp(prompt string) {
  req := gpt3.CompletionRequest{
    Prompt: []string{prompt},
    MaxTokens: gpt3.IntPtr(1000),
    Temperature: gpt3.Float32Ptr(0),
  }

  err := ai.Cli.CompletionStreamWithEngine(
    ai.Ctx, gpt3.TextDavinvi003Engine, req,
    func(resp *gpt3.CompletionResponse) {
      fmt.Print(resp.Choices[0].Text)
    },
    )

  if err != nil {
    panic(err)
  }

  fmt.Println("")
}

