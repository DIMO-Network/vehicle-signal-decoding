INSERT INTO templates (template_name, parent_template_name, template_type, version)
    VALUES ('default-ice-can11', null, null, 'v1.0.0');

INSERT INTO pid_configs (template_name, pid, formula, interval_seconds, protocol, signal_name)
VALUES ('default-ice-can11', E'\\x30', '31|8@0+ (1,0) [0|255] "count"', 0, 'CAN11_500', 'warmups_since_dtc_clear');
INSERT INTO pid_configs (template_name, pid, formula, interval_seconds, protocol, signal_name)
VALUES ('default-ice-can11', E'\\x31', '31|16@0+ (1,0) [0|65535] "km" ', 0, 'CAN11_500', 'distance_since_dtc_clear');
INSERT INTO pid_configs (template_name, pid, formula, interval_seconds, protocol, signal_name)
VALUES ('default-ice-can11', E'\\x21', '31|16@0+ (1,0) [0|65535] "km"', 0, 'CAN11_500', 'distance_w_mil');
INSERT INTO pid_configs (template_name, pid, formula, interval_seconds, protocol, signal_name)
VALUES ('default-ice-can11', E'\\x2c', '31|8@0+ (0.392156862745098,0) [0|100] "%"', 20, 'CAN11_500', 'commanded_egr');
INSERT INTO pid_configs (template_name, pid, formula, interval_seconds, protocol, signal_name)
VALUES ('default-ice-can11', E'\\x1f', '31|16@0+ (1,0) [0|65535] "seconds"', 30, 'CAN11_500', 'run_time');
INSERT INTO pid_configs (template_name, pid, formula, interval_seconds, protocol, signal_name)
VALUES ('default-ice-can11', E'\\x06', '31|8@0+ (0.78125,-100) [-100|99.21875] "%"', 10, 'CAN11_500', 'short_fuel_trim');
INSERT INTO pid_configs (template_name, pid, formula, interval_seconds, protocol, signal_name)
VALUES ('default-ice-can11', E'\\x11', '31|8@0+ (0.39216,0) [0|100] "%"', 5, 'CAN11_500', 'throttle_position');
INSERT INTO pid_configs (template_name, pid, formula, interval_seconds, protocol, signal_name)
VALUES ('default-ice-can11', E'\\x07', '31|8@0+ (0.78125,-100) [-100|99.21875] "%"', 60, 'CAN11_500', 'long_fuel_trim');
INSERT INTO pid_configs (template_name, pid, formula, interval_seconds, protocol, signal_name)
VALUES ('default-ice-can11', E'\\x0b', '31|8@0+ (1,0) [0|255] "kPa"', 10, 'CAN11_500', 'intake_pressure');
INSERT INTO pid_configs (template_name, pid, formula, interval_seconds, protocol, signal_name)
VALUES ('default-ice-can11', E'\\x10', '31|16@0+ (0.01,0) [0|655.35] "grams/sec"', 10, 'CAN11_500', 'maf');
INSERT INTO pid_configs (template_name, pid, formula, interval_seconds, protocol, signal_name)
VALUES ('default-ice-can11', E'\\x33', '31|8@0+ (1,0) [0|255] "kPa"', 60, 'CAN11_500', 'barometric_pressure');
INSERT INTO pid_configs (template_name, pid, formula, interval_seconds, protocol, signal_name)
VALUES ('default-ice-can11', E'\\x0d', '31|8@0+ (1,0) [0|255] "km/h"', 10, 'CAN11_500', 'vehicle_speed');
INSERT INTO pid_configs (template_name, pid, formula, interval_seconds, protocol, signal_name)
VALUES ('default-ice-can11', E'\\x04', '31|8@0+ (0.39216,0) [0|100] "%"', 10, 'CAN11_500', 'engine_load');
INSERT INTO pid_configs (template_name, pid, formula, interval_seconds, protocol, signal_name)
VALUES ('default-ice-can11', E'\\x46', '31|8@0+ (1,-40) [-40|215] "Celcius"', 60, 'CAN11_500', 'ambient_air_temp');
INSERT INTO pid_configs (template_name, pid, formula, interval_seconds, protocol, signal_name)
VALUES ('default-ice-can11', E'\\x0f', '31|8@0+ (1,-40) [-40|215] "Celcius"', 10, 'CAN11_500', 'intake_temp');
INSERT INTO pid_configs (template_name, pid, formula, interval_seconds, protocol, signal_name)
VALUES ('default-ice-can11', E'\\x2f', '31|8@0+ (0.392156862745098,0) [0|100] "%"', 60, 'CAN11_500', 'fuel_level');
INSERT INTO pid_configs (template_name, pid, formula, interval_seconds, protocol, signal_name)
VALUES ('default-ice-can11', E'\\x0c', '31|16@0+ (0.25,0) [0|16383.75] "rpm"', 5, 'CAN11_500', 'rpm');
INSERT INTO pid_configs (template_name, pid, formula, interval_seconds, protocol, signal_name)
VALUES ('default-ice-can11', E'\\x67', '31|8@0+ (1,0) [0|255] "Celcius"', 30, 'CAN11_500', 'coolant_temp');

-- 2019+ with odometer
INSERT INTO templates (template_name, parent_template_name, template_type, version)
VALUES ('2019plus-ice-can11', 'default-ice-can11', null, 'v1.0.0');

INSERT INTO pid_configs (template_name, pid, formula, interval_seconds, protocol, signal_name)
VALUES ('2019plus-ice-can11', E'\\xa6', '31|32@0+ (0.1,0) [1|429496730] "km"', 60, 'CAN11_500', 'odometer');

--- protocol 7 / can 29

INSERT INTO templates (template_name, parent_template_name, template_type, version)
VALUES ('default-ice-can29', null, null, 'v1.0.0');

INSERT INTO pid_configs (header, template_name, pid, formula, interval_seconds, protocol, signal_name)
VALUES ('18DB33F1', 'default-ice-can29', E'\\x30', '31|8@0+ (1,0) [0|255] "count"', 0, 'CAN11_500', 'warmups_since_dtc_clear');
INSERT INTO pid_configs (header, template_name, pid, formula, interval_seconds, protocol, signal_name)
VALUES ('18DB33F1', 'default-ice-can29', E'\\x31', '31|16@0+ (1,0) [0|65535] "km" ', 0, 'CAN11_500', 'distance_since_dtc_clear');
INSERT INTO pid_configs (header, template_name, pid, formula, interval_seconds, protocol, signal_name)
VALUES ('18DB33F1', 'default-ice-can29', E'\\x21', '31|16@0+ (1,0) [0|65535] "km"', 0, 'CAN11_500', 'distance_w_mil');
INSERT INTO pid_configs (header, template_name, pid, formula, interval_seconds, protocol, signal_name)
VALUES ('18DB33F1', 'default-ice-can29', E'\\x2c', '31|8@0+ (0.392156862745098,0) [0|100] "%"', 20, 'CAN11_500', 'commanded_egr');
INSERT INTO pid_configs (header, template_name, pid, formula, interval_seconds, protocol, signal_name)
VALUES ('18DB33F1', 'default-ice-can29', E'\\x1f', '31|16@0+ (1,0) [0|65535] "seconds"', 30, 'CAN11_500', 'run_time');
INSERT INTO pid_configs (header, template_name, pid, formula, interval_seconds, protocol, signal_name)
VALUES ('18DB33F1', 'default-ice-can29', E'\\x06', '31|8@0+ (0.78125,-100) [-100|99.21875] "%"', 10, 'CAN11_500', 'short_fuel_trim');
INSERT INTO pid_configs (header, template_name, pid, formula, interval_seconds, protocol, signal_name)
VALUES ('18DB33F1', 'default-ice-can29', E'\\x11', '31|8@0+ (0.39216,0) [0|100] "%"', 5, 'CAN11_500', 'throttle_position');
INSERT INTO pid_configs (header, template_name, pid, formula, interval_seconds, protocol, signal_name)
VALUES ('18DB33F1', 'default-ice-can29', E'\\x07', '31|8@0+ (0.78125,-100) [-100|99.21875] "%"', 60, 'CAN11_500', 'long_fuel_trim');
INSERT INTO pid_configs (header, template_name, pid, formula, interval_seconds, protocol, signal_name)
VALUES ('18DB33F1', 'default-ice-can29', E'\\x0b', '31|8@0+ (1,0) [0|255] "kPa"', 10, 'CAN11_500', 'intake_pressure');
INSERT INTO pid_configs (header, template_name, pid, formula, interval_seconds, protocol, signal_name)
VALUES ('18DB33F1', 'default-ice-can29', E'\\x10', '31|16@0+ (0.01,0) [0|655.35] "grams/sec"', 10, 'CAN11_500', 'maf');
INSERT INTO pid_configs (header, template_name, pid, formula, interval_seconds, protocol, signal_name)
VALUES ('18DB33F1', 'default-ice-can29', E'\\x33', '31|8@0+ (1,0) [0|255] "kPa"', 60, 'CAN11_500', 'barometric_pressure');
INSERT INTO pid_configs (header, template_name, pid, formula, interval_seconds, protocol, signal_name)
VALUES ('18DB33F1', 'default-ice-can29', E'\\x0d', '31|8@0+ (1,0) [0|255] "km/h"', 10, 'CAN11_500', 'vehicle_speed');
INSERT INTO pid_configs (header, template_name, pid, formula, interval_seconds, protocol, signal_name)
VALUES ('18DB33F1', 'default-ice-can29', E'\\x04', '31|8@0+ (0.39216,0) [0|100] "%"', 10, 'CAN11_500', 'engine_load');
INSERT INTO pid_configs (header, template_name, pid, formula, interval_seconds, protocol, signal_name)
VALUES ('18DB33F1', 'default-ice-can29', E'\\x46', '31|8@0+ (1,-40) [-40|215] "Celcius"', 60, 'CAN11_500', 'ambient_air_temp');
INSERT INTO pid_configs (header, template_name, pid, formula, interval_seconds, protocol, signal_name)
VALUES ('18DB33F1', 'default-ice-can29', E'\\x0f', '31|8@0+ (1,-40) [-40|215] "Celcius"', 10, 'CAN11_500', 'intake_temp');
INSERT INTO pid_configs (header, template_name, pid, formula, interval_seconds, protocol, signal_name)
VALUES ('18DB33F1', 'default-ice-can29', E'\\x2f', '31|8@0+ (0.392156862745098,0) [0|100] "%"', 60, 'CAN11_500', 'fuel_level');
INSERT INTO pid_configs (header, template_name, pid, formula, interval_seconds, protocol, signal_name)
VALUES ('18DB33F1', 'default-ice-can29', E'\\x0c', '31|16@0+ (0.25,0) [0|16383.75] "rpm"', 5, 'CAN11_500', 'rpm');
INSERT INTO pid_configs (header, template_name, pid, formula, interval_seconds, protocol, signal_name)
VALUES ('18DB33F1', 'default-ice-can29', E'\\x67', '31|8@0+ (1,0) [0|255] "Celcius"', 30, 'CAN11_500', 'coolant_temp');

-- 2019+ odometer
INSERT INTO templates (template_name, parent_template_name, template_type, version)
VALUES ('2019plus-ice-can29', 'default-ice-can29', null, 'v1.0.0');

INSERT INTO pid_configs (header, template_name, pid, formula, interval_seconds, protocol, signal_name)
VALUES ('18DB33F1', '2019plus-ice-can29', E'\\xa6', '31|32@0+ (0.1,0) [1|429496730] "km"', 60, 'CAN11_500', 'odometer');
