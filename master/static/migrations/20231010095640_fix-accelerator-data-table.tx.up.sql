ALTER TABLE public.allocation_accelerators
    DROP CONSTRAINT allocation_accelerators_pkey,
    ADD id INT GENERATED BY DEFAULT AS IDENTITY,
    ADD CONSTRAINT allocation_accelerators_pkey PRIMARY KEY(id);
