CREATE TABLE log_activities (
    id uuid PRIMARY KEY,
    action_type varchar(100) NOT NULL,
    ip_address inet,
    duration_ms integer,
    metadata jsonb,
    created_at timestamptz NOT NULL DEFAULT now(),
    user_id uuid
);