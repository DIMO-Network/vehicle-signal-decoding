INSERT INTO templates (template_name, parent_template_name, template_type, version)
    VALUES ('default-ice-can11', null, null, 'v1.0.0');

INSERT INTO pid_configs (template_name, pid, formula, interval_seconds, protocol, signal_name)
VALUES ('default-ice-can11', E'\\x30', '-', 0, 'CAN11_500', 'warmups_since_dtc_clear');
INSERT INTO pid_configs (template_name, pid, formula, interval_seconds, protocol, signal_name)
VALUES ('default-ice-can11', E'\\x31', '-', 0, 'CAN11_500', 'distance_since_dtc_clear');
INSERT INTO pid_configs (template_name, pid, formula, interval_seconds, protocol, signal_name)
VALUES ('default-ice-can11', E'\\x21', '-', 0, 'CAN11_500', 'distance_w_mil');
INSERT INTO pid_configs (template_name, pid, formula, interval_seconds, protocol, signal_name)
VALUES ('default-ice-can11', E'\\x2c', '-', 20, 'CAN11_500', 'commanded_egr');
INSERT INTO pid_configs (template_name, pid, formula, interval_seconds, protocol, signal_name)
VALUES ('default-ice-can11', E'\\x1f', '-', 10, 'CAN11_500', 'run_time');
INSERT INTO pid_configs (template_name, pid, formula, interval_seconds, protocol, signal_name)
VALUES ('default-ice-can11', E'\\x04', '-', 10, 'CAN11_500', 'engine_load');
INSERT INTO pid_configs (template_name, pid, formula, interval_seconds, protocol, signal_name)
VALUES ('default-ice-can11', E'\\x06', '-', 10, 'CAN11_500', 'short_fuel_trim');
INSERT INTO pid_configs (template_name, pid, formula, interval_seconds, protocol, signal_name)
VALUES ('default-ice-can11', E'\\x46', '-', 60, 'CAN11_500', 'ambient_air_temp');
INSERT INTO pid_configs (template_name, pid, formula, interval_seconds, protocol, signal_name)
VALUES ('default-ice-can11', E'\\x11', '-', 5, 'CAN11_500', 'throttle_position');
INSERT INTO pid_configs (template_name, pid, formula, interval_seconds, protocol, signal_name)
VALUES ('default-ice-can11', E'\\x07', '-', 60, 'CAN11_500', 'long_fuel_trim');
INSERT INTO pid_configs (template_name, pid, formula, interval_seconds, protocol, signal_name)
VALUES ('default-ice-can11', E'\\x0b', '-', 10, 'CAN11_500', 'intake_pressure');
INSERT INTO pid_configs (template_name, pid, formula, interval_seconds, protocol, signal_name)
VALUES ('default-ice-can11', E'\\x10', 'decimal', 10, 'CAN11_500', 'maf');
INSERT INTO pid_configs (template_name, pid, formula, interval_seconds, protocol, signal_name)
VALUES ('default-ice-can11', E'\\x33', '-', 60, 'CAN11_500', 'barometric_pressure');
INSERT INTO pid_configs (template_name, pid, formula, interval_seconds, protocol, signal_name)
VALUES ('default-ice-can11', E'\\x0d', '-', 10, 'CAN11_500', 'vehicle_speed');
INSERT INTO pid_configs (template_name, pid, formula, interval_seconds, protocol, signal_name)
VALUES ('default-ice-can11', E'\\x04', '-', 10, 'CAN11_500', 'engine_load');
INSERT INTO pid_configs (template_name, pid, formula, interval_seconds, protocol, signal_name)
VALUES ('default-ice-can11', E'\\x46', '-', 10, 'CAN11_500', 'ambient_air_temp');
INSERT INTO pid_configs (template_name, pid, formula, interval_seconds, protocol, signal_name)
VALUES ('default-ice-can11', E'\\x0f', '-', 10, 'CAN11_500', 'intake_temp');
INSERT INTO pid_configs (template_name, pid, formula, interval_seconds, protocol, signal_name)
VALUES ('default-ice-can11', E'\\x2f', '-', 10, 'CAN11_500', 'fuel_level');
INSERT INTO pid_configs (template_name, pid, formula, interval_seconds, protocol, signal_name)
VALUES ('default-ice-can11', E'\\x0c', '-', 10, 'CAN11_500', 'rpm');
INSERT INTO pid_configs (template_name, pid, formula, interval_seconds, protocol, signal_name)
VALUES ('default-ice-can11', E'\\x67', '-', 10, 'CAN11_500', 'coolant_temp');

--- protocol 7 / can 29

INSERT INTO templates (template_name, parent_template_name, template_type, version)
VALUES ('default-ice-can29', null, null, 'v1.0.0');

