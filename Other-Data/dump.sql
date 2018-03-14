--
-- PostgreSQL database dump
--

-- Dumped from database version 9.5.8
-- Dumped by pg_dump version 9.5.8

SET statement_timeout = 0;
SET lock_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: movie; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE movie WITH TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'en_IN' LC_CTYPE = 'en_IN';


ALTER DATABASE movie OWNER TO postgres;

\connect movie

SET statement_timeout = 0;
SET lock_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET search_path = public, pg_catalog;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: record; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE record (
    id character varying(30),
    title text,
    release_year text,
    rating double precision,
    genres text[]
);


ALTER TABLE record OWNER TO postgres;

--
-- Data for Name: record; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY record (id, title, release_year, rating, genres) FROM stdin;
tt0090191	The Avengers	2015	8.09999999999999964	{action,drama,comedy,horror}
tt0090189	C company	1984	6.5	{action,drama,comedy}
tt0090182	3 idiot	2013	7.70000000000000018	{drama,comedy}
tt0090181	YJHD	2011	9.19999999999999929	{romance,drama}
tt7329858	Pari	2018	6.90000000000000036	{"Horror,",Mystery}
tt7218518	Padman	2018	8.5	{"Biography,","Comedy,",Drama}
tt5935704	Padmaavat	2018	7.29999999999999982	{"Drama,","History,",Romance}
tt6774212	Aiyaary	2018	5.79999999999999982	{"Action,","Crime,",Drama}
tt0090190	The Toxic Avenger	2000	6.29999999999999982	{"drama,action"}
\.


--
-- Name: record_id_uindex; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX record_id_uindex ON record USING btree (id);


--
-- Name: public; Type: ACL; Schema: -; Owner: postgres
--

REVOKE ALL ON SCHEMA public FROM PUBLIC;
REVOKE ALL ON SCHEMA public FROM postgres;
GRANT ALL ON SCHEMA public TO postgres;
GRANT ALL ON SCHEMA public TO PUBLIC;


--
-- PostgreSQL database dump complete
--

