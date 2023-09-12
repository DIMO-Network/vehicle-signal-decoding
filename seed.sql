INSERT INTO templates (template_name, parent_template_name, template_type, version)
    VALUES ('default-ice-can11', null, null, 'v1.0.0');

INSERT INTO pid_configs (id, template_name, pid, formula, interval_seconds, protocol, signal_name)
VALUES (7, 'default-ice-can11', E'\\x1f', '-', 10, 'CAN11_500', 'run_time');
INSERT INTO pid_configs (id, template_name, pid, formula, interval_seconds, protocol, signal_name)
VALUES (11, 'default-ice-can11', E'\\x04', '-', 10, 'CAN11_500', 'engine_load');
INSERT INTO pid_configs (id, template_name, pid, formula, interval_seconds, protocol, signal_name)
VALUES (5, 'default-ice-can11', E'\\x06', '-', 10, 'CAN11_500', 'short_fuel_trim');
INSERT INTO pid_configs (id, template_name, pid, formula, interval_seconds, protocol, signal_name)
VALUES (12, 'default-ice-can11', E'\\x46', '-', 60, 'CAN11_500', 'ambient_air_temp');
INSERT INTO pid_configs (id, template_name, pid, formula, interval_seconds, protocol, signal_name)
VALUES (9, 'default-ice-can11', E'\\x11', '-', 5, 'CAN11_500', 'throttle_position');
INSERT INTO pid_configs (id, template_name, pid, formula, interval_seconds, protocol, signal_name)
VALUES (6, 'default-ice-can11', E'\\x07', '-', 60, 'CAN11_500', 'long_fuel_trim');
INSERT INTO pid_configs (id, template_name, pid, formula, interval_seconds, protocol, signal_name)
VALUES (4, 'default-ice-can11', E'\\x0b', '-', 10, 'CAN11_500', 'intake_pressure');
INSERT INTO pid_configs (id, template_name, pid, formula, interval_seconds, protocol, signal_name)
VALUES (2, 'default-ice-can11', E'\\x10', 'decimal', 10, 'CAN11_500', 'maf');
INSERT INTO pid_configs (id, template_name, pid, formula, interval_seconds, protocol, signal_name)
VALUES (8, 'default-ice-can11', E'\\x33', '-', 60, 'CAN11_500', 'barometric_pressure');
INSERT INTO pid_configs (id, template_name, pid, formula, interval_seconds, protocol, signal_name)
VALUES (10, 'default-ice-can11', E'\\x0d', '-', 10, 'CAN11_500', 'vehicle_speed');

