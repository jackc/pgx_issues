create procedure testproc(inout foo int) language plpgsql as $$
begin
  foo := 42;
end;
$$;
