package main

import (
	"context"
	"github.com/DIMO-Network/model-garage/pkg/schema"
	"github.com/DIMO-Network/shared/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/config"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	"github.com/rs/zerolog"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"gopkg.in/yaml.v3"
	"strings"
)

func SyncCovesaSignalNames(ctx context.Context, logger zerolog.Logger, config *config.Settings, args []string) {
	sqlDb := db.NewDbConnectionFromSettings(ctx, &s.DB, true)

	sqlDb.WaitForDB(logger)

	pidConfigs, err := models.PidConfigs().All(ctx, sqlDb.DBS().Reader)

	if err != nil {
		logger.Fatal().Err(err).Msg("failed to get pid configs")
		return
	}

	definitions := []byte(schema.DefinitionsYAML())

	type convertibleSignal struct {
		VspecName   string `yaml:"vspecName"`
		IsArray     bool   `yaml:"isArray"`
		Conversions []struct {
			OriginalName string `yaml:"originalName"`
			OriginalType string `yaml:"originalType"`
			IsArray      bool   `yaml:"isArray"`
		}
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
				if strings.EqualFold(pidConfig.SignalName, conversion.OriginalName) {

					pidConfig.VSSCovesaName = null.StringFrom(convertibleSignal.VspecName)
					_, err := pidConfig.Update(ctx, sqlDb.DBS().Writer, boil.Infer())

					if err != nil {
						logger.Fatal().Err(err).Msg("failed to update pid config")
						return
					}

					totalUpdated++

				} else {

					if alreadyInSlice(unTranslatedSignals, pidConfig.SignalName) {
						continue
					}

					unTranslatedSignals = append(unTranslatedSignals, pidConfig.SignalName)
				}
			}
		}
	}

	if len(unTranslatedSignals) > 0 {
		logger.Info().
			Int("total_updated", totalUpdated)

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
