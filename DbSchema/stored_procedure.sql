DROP FUNCTION IF EXISTS GET_LIST_OF_PEOPLE;
CREATE FUNCTION GET_LIST_OF_PEOPLE(_search_term VARCHAR)
RETURNS 
    TABLE(
        id BIGINT,
        name VARCHAR,
        date_of_birth TIMESTAMP
    )
LANGUAGE plpgsql
as
$$
BEGIN
    RETURN QUERY SELECT p.id, p.name, p.date_of_birth FROM people as p WHERE p.name LIKE '%' || _search_term || '%';
END;
$$;


DROP FUNCTION IF EXISTS ADD_PEOPLE;
CREATE FUNCTION ADD_PEOPLE(
        _name VARCHAR, 
        _date_of_birth TIMESTAMP 
    )
RETURNS INT
LANGUAGE plpgsql
as
$$
DECLARE
    return_id INT;
BEGIN
    INSERT INTO people(name, date_of_birth) VALUES (_name, _date_of_birth) RETURNING id INTO return_id;
    RETURN return_id;
END;
$$;

