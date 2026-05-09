ALTER TABLE public.posts
ADD COLUMN IF NOT EXISTS university_id UUID REFERENCES public.universities(id) ON DELETE SET NULL,
ADD COLUMN IF NOT EXISTS is_published BOOLEAN NOT NULL DEFAULT FALSE;

UPDATE public.posts
SET is_published = TRUE
WHERE published_at IS NOT NULL;
