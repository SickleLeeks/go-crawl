package work

import (
	"crawl/LearnGo-crawl/engine"
	"crawl/LearnGo-crawl/parse/zhenai"
	"errors"
	"fmt"
	"log"
)

type SerializeParse struct {
	Name string
	Args interface{}
}
type Request struct {
	Url   string
	Parse SerializeParse
}
type ParseResult struct {
	items    []engine.Item
	Requests []Request
}

func SerializeResult(r engine.ParseResult) ParseResult {
	result := ParseResult{items: r.Items}
	for _, req := range r.Requests {
		result.Requests = append(result.Requests, SerializeRequest(req))
	}
	return result
}
func SerializeRequest(r engine.Request) Request {
	name, args := r.Parse.Serialize()
	return Request{
		Url: r.Url,
		Parse: SerializeParse{
			Name: name,
			Args: args,
		},
	}
}
func DeserializeResult(r ParseResult) engine.ParseResult {
	result := engine.ParseResult{
		Items: r.items,
	}
	for _, req := range r.Requests {
		engineReq, err := DeserializeRequest(req)
		if err != nil {
			log.Printf("error deserialized: %v", err)
			continue
		}
		result.Requests = append(result.Requests, engineReq)
	}
	return result
}
func DeserializeRequest(r Request) (engine.Request, error) {
	parse, err := deserializeParse(r.Parse)
	if err != nil {
		return engine.Request{}, err
	}
	return engine.Request{
		Url:   r.Url,
		Parse: parse,
	}, nil
}

func deserializeParse(p SerializeParse) (engine.Parser, error) {
	switch p.Name {
	case "ParseCitylist":
		return engine.NewFuncparse(zhenai.ParseCityList, "ParseCitylist"), nil
	case "Parsecity":
		return engine.NewFuncparse(zhenai.ParseCity, "Parsecity"), nil
	case "ProfileParse":
		if userName, ok := p.Args.(string); ok {
			return zhenai.NewprofileParse(userName), nil
		} else {
			return nil, fmt.Errorf("invilid args: %v", p.Args)
		}
	case "Nilparse":
		return engine.Nilparse{}, nil
	default:
		return nil, errors.New("Unknown parse name")
	}
}
