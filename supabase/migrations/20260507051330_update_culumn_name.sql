ALTER TABLE public.universities
RENAME COLUMN is_public TO is_published;

ALTER TABLE public.faculties
RENAME COLUMN is_public TO is_published;