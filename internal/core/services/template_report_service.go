package services

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/config"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"net/http"
	"os"
)

type TemplateReportService struct {
	log *zerolog.Logger
}

// todo call this from command line. fixup all log statements. get credentials from settings. consider storing in memory
// note that may need to change scope of the credentials like below.

// BuildReport uses google api to generate a google spreadsheet
func BuildReport(settings *config.Settings) error {
	ctx := context.Background()

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON([]byte(settings.GoogleSheetsCreds), sheets.SpreadsheetsScope)
	if err != nil {
		return errors.Wrap(err, "Unable to parse client secret file to config")
	}
	client, err := getClient(config)

	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return errors.Wrap(err, "Unable to retrieve Sheets client")
	}

	spreadsheetId := "1dKcJQXX0WazF4zjxr-Ziul7lHfKU1f20KmE7qEMSepA" // in Vehicle Decoding - Templates Report

	// Add new rows
	writeRange := "Sheet1!A1"
	var vr sheets.ValueRange
	myval := []interface{}{"Value1", "Value2", "Value3"}
	vr.Values = append(vr.Values, myval)
	_, err = srv.Spreadsheets.Values.Append(spreadsheetId, writeRange, &vr).ValueInputOption("RAW").Do()
	if err != nil {
		return errors.Wrap(err, "Unable to add new value to Spreadsheet")
	}

	// Add new columns
	updateRange := "Sheet1!D1"
	var ur sheets.ValueRange
	myColumnValues := []interface{}{"Column1", "Column2", "Column3"}
	ur.Values = append(ur.Values, myColumnValues)
	_, err = srv.Spreadsheets.Values.Update(spreadsheetId, updateRange, &ur).ValueInputOption("RAW").Do()
	if err != nil {
		return errors.Wrap(err, "Unable to add new value to Spreadsheet")
	}

	fmt.Println("Rows and columns added successfully.")

	return nil
}

// getClient uses a Context and Config to retrieve a Token
// then generate a Client. It returns the generated Client.
func getClient(config *oauth2.Config) (*http.Client, error) {
	tokFile := "token.json"
	// grab existing token that is stored locally - replace this with in memory storage
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok, err = getTokenFromWeb(config)
		if err != nil {
			return nil, err
		}
		err = saveToken(tokFile, tok)
		if err != nil {
			return nil, err
		}
	}
	return config.Client(context.Background(), tok), nil
}

// tokenFromFile retrieves a Token from a given file path.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// getTokenFromWeb uses Config to request a Token.
// It returns the retrieved Token.
func getTokenFromWeb(config *oauth2.Config) (*oauth2.Token, error) {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		return nil, errors.Wrap(err, "Unable to read authorization code")
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to retrieve token from web")
	}
	return tok, nil
}

// saveToken saves a Token to a file path.
func saveToken(path string, token *oauth2.Token) error {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.Create(path)
	if err != nil {
		return errors.Wrap(err, "Unable to cache oauth token")
	}
	defer f.Close()
	return json.NewEncoder(f).Encode(token)
}
