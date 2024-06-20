UPDATE person SET marital_status = CASE 
    WHEN is_married = true THEN 'married'::marital_status
    ELSE 'not married'::marital_status
END;