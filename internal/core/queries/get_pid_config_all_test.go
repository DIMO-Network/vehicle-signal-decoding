package queries

import (
	"context"
	b_rand "crypto/rand"
	"math/rand"
	"testing"

	"github.com/DIMO-Network/shared/pkg/db"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/db/models"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/infrastructure/dbtest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

const dbSchemaName = "vehicle_signal_decoding"

func TestGetPidsByTemplate_HandleDuplicates(t *testing.T) {
	// this test proves get the child pid if duplicated signal name with parent
	ctx := context.Background()
	pdb, container := dbtest.StartContainerDatabase(ctx, dbSchemaName, t, migrationsDirRelPath)
	//logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	const templateName = "toyota-tnga"
	// add a few templates
	parentTmpl := models.Template{
		TemplateName: "default-parent1",
		Version:      "1.0",
		Protocol:     models.CanProtocolTypeCAN11_500,
		Powertrain:   "ICE",
	}
	err := parentTmpl.Insert(ctx, pdb.DBS().Writer, boil.Infer())
	require.NoError(t, err)

	childTmpl1 := models.Template{
		TemplateName:       templateName,
		ParentTemplateName: null.StringFrom(parentTmpl.TemplateName),
		Version:            "1.0",
		Protocol:           models.CanProtocolTypeCAN11_500,
		Powertrain:         "ICE",
	}
	err = childTmpl1.Insert(ctx, pdb.DBS().Writer, boil.Infer())
	require.NoError(t, err)
	// add 3 pids
	createPID(ctx, t, parentTmpl, "soc", "f1", pdb)
	createPID(ctx, t, parentTmpl, "odometer", "f1", pdb)
	createPID(ctx, t, parentTmpl, "tires.left", "f1", pdb)
	createPID(ctx, t, childTmpl1, "odometer", "f2", pdb)
	createPID(ctx, t, childTmpl1, "tires.left", "f2", pdb)

	gotPids, gotTemplate, err := GetPidsByTemplate(ctx, pdb.DBS, &GetPidsQueryRequest{TemplateName: templateName})
	require.NoError(t, err)

	assert.Equal(t, templateName, gotTemplate.TemplateName)
	assert.Len(t, gotPids, 3)
	for _, pid := range gotPids {
		if pid.SignalName == "odometer" {
			assert.Equal(t, "f2", pid.Formula)
		}
		if pid.SignalName == "tires.left" {
			assert.Equal(t, "f2", pid.Formula)
		}
	}

	if err := container.Terminate(ctx); err != nil {
		t.Fatal(err)
	}
}

func createPID(ctx context.Context, t *testing.T, parentTmpl models.Template, signal, formula string, pdb db.Store) {
	var bytes [2]byte
	// Generate random 2 bytes
	b_rand.Read(bytes[:]) //nolint

	pid1 := models.PidConfig{
		ID:              rand.Int63(),
		SignalName:      signal,
		TemplateName:    parentTmpl.TemplateName,
		Header:          bytes[:],
		Mode:            bytes[1:],
		Pid:             bytes[0:],
		Formula:         formula,
		IntervalSeconds: 60,
		Protocol:        null.StringFrom(models.CanProtocolTypeCAN11_500),
	}
	err := pid1.Insert(ctx, pdb.DBS().Writer, boil.Infer())
	require.NoError(t, err)
}
