-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
SET search_path = vehicle_signal_decoding_api, public;

DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM information_schema.columns 
        WHERE table_schema = 'vehicle_signal_decoding_api' 
        AND table_name = 'test_signals' 
        AND column_name = 'signals'
    ) THEN
        ALTER TABLE test_signals ADD COLUMN signals jsonb;
    END IF;
END $$;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
SET search_path = vehicle_signal_decoding_api, public;
alter table test_signals drop column signals;
-- +goose StatementEnd
