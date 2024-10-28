-- +goose Up
-- +goose StatementBegin
CREATE TABLE patient_scans(
  id serial primary key,
	patient_id int references patients (id) on delete set null
);

INSERT INTO patient_scans (patient_id)
VALUES (1);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE patient_scans;
-- +goose StatementEnd
