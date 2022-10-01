create extension hstore;

create table test
(
	id bigserial not null primary key,
	changes hstore
);
INSERT INTO test (id, changes) VALUES (1, 'appointment_time => null, created => "2016-07-21 15:15:23.187727+03", postponed_till => null, latitude => 36.2389755249023, created_local => "2016-07-21 15:15:23+03", ticket_id => null, time_slots => "[{\"start_time\": \"2016-07-24T00:00:00\", \"end_time\": \"2016-07-24T23:59:00\"}, {\"start_time\": \"2016-07-25T00:00:00\", \"end_time\": \"2016-07-25T23:59:00\"}, {\"start_time\": \"2016-07-26T00:00:00\", \"end_time\": \"2016-07-26T23:59:00\"}, {\"start_time\": \"2016-07-27T00:00:00\", \"end_time\": \"2016-07-27T23:59:00\"}, {\"start_time\": \"2016-07-28T00:00:00\", \"end_time\": \"2016-07-28T23:59:00\"}, {\"start_time\": \"2016-07-22T00:00:00\", \"end_time\": \"2016-07-22T23:59:00\"}, {\"start_time\": \"2016-07-23T00:00:00\", \"end_time\": \"2016-07-23T23:59:00\"}]", patient_id => 1, specialization_code => therapist, status_updated => "2016-07-21 15:15:23.285021+03", id => 1, doctor_name => null, cancelation_reason_id => null, clinic_id => 2, status => pending, longitude => 99.008171081543');
