package main

import (
	"context"
	"strings"

	nativestatus "github.com/DIMO-Network/model-garage/pkg/nativestatus/schema"
	vssschema "github.com/DIMO-Network/model-garage/pkg/schema"
	"github.com/DIMO-Network/shared/pkg/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/config"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	"github.com/aarondl/null/v8"
	"github.com/aarondl/sqlboiler/v4/boil"
	"github.com/rs/zerolog"
	"gopkg.in/yaml.v3"
)

// nolint
type convertibleSignal struct {
	VspecName   string `yaml:"vspecName"`
	IsArray     bool   `yaml:"isArray"`
	Conversions []struct {
		OriginalName string `yaml:"originalName"`
		OriginalType string `yaml:"originalType"`
		IsArray      bool   `yaml:"isArray"`
	}
}

func SyncCovesaSignalNames(ctx context.Context, logger zerolog.Logger, config *config.Settings, _ []string) {
	sqlDb := db.NewDbConnectionFromSettings(ctx, &config.DB, true)

	sqlDb.WaitForDB(logger)

	pidConfigs, err := models.PidConfigs().All(ctx, sqlDb.DBS().Reader)

	if err != nil {
		logger.Fatal().Err(err).Msg("failed to get pid configs")
		return
	}

	definitions := []byte(nativestatus.DefinitionsYAML())

	covesaSignals := strings.NewReader(vssschema.VssRel42DIMO())
	signals, err := vssschema.LoadSignalsCSV(covesaSignals)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to load covesa signals")
		return
	}

	convertibleSignals := make([]convertibleSignal, 0)

	err = yaml.Unmarshal(definitions, &convertibleSignals)

	if err != nil {
		logger.Fatal().Err(err).Msg("failed to unmarshal pid signal translations")
		return
	}

	unTranslatedSignals := make([]string, 0)
	totalUpdated := 0

	for _, pidConfig := range pidConfigs {

		if pidConfig.VSSCovesaName.Valid {
			continue
		}

		for _, convertibleSignal := range convertibleSignals {
			for _, conversion := range convertibleSignal.Conversions {

				pidSignalName := strings.ReplaceAll(pidConfig.SignalName, " ", "")
				conversionOriginalName := strings.ReplaceAll(conversion.OriginalName, " ", "")

				if strings.EqualFold(pidSignalName, conversionOriginalName) {

					vspecName := strings.ReplaceAll(convertibleSignal.VspecName, " ", "")

					pidConfig.VSSCovesaName = null.StringFrom(vspecName)

					for _, signal := range signals {
						if strings.EqualFold(signal.Name, vspecName) {
							pidConfig.Unit = null.StringFrom(signal.Unit)
							break
						}
					}

					_, err := pidConfig.Update(ctx, sqlDb.DBS().Writer, boil.Infer())

					if err != nil {
						logger.Fatal().Err(err).Msg("failed to update pid config")
						return
					}

					totalUpdated++

				} else {

					if alreadyInSlice(unTranslatedSignals, pidSignalName) {
						continue
					}

					unTranslatedSignals = append(unTranslatedSignals, pidSignalName)
				}
			}
		}
	}

	if len(unTranslatedSignals) > 0 {
		logger.Info().
			Int("total_updated", totalUpdated).
			Msg("successfully updated pid configs")

		logger.Warn().
			Strs("untranslated_signals", unTranslatedSignals).
			Msg("some signals were not translated")

	}

	logger.Info().Msg("successfully updated pid configs")
}

func alreadyInSlice(slice []string, val string) bool {
	for _, item := range slice {
		if strings.EqualFold(item, val) {
			return true
		}
	}
	return false
}
