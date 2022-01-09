--
-- PostgreSQL database dump
--

-- Dumped from database version 14.1
-- Dumped by pg_dump version 14.1

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: owned_titles; Type: TABLE; Schema: public; Owner: wiisoap
--

CREATE TABLE public.owned_titles (
                                     account_id integer NOT NULL,
                                     title_id character varying(16) NOT NULL,
                                     version integer
);


ALTER TABLE public.owned_titles OWNER TO wiisoap;

--
-- Name: tickets; Type: TABLE; Schema: public; Owner: wiisoap
--

CREATE TABLE public.tickets (
                                title_id character varying(16) NOT NULL,
                                ticket bytea,
                                version integer
);


ALTER TABLE public.tickets OWNER TO wiisoap;

--
-- Name: userbase; Type: TABLE; Schema: public; Owner: wiisoap
--

CREATE TABLE public.userbase (
                                 device_id bigint NOT NULL,
                                 device_token character varying(21) NOT NULL,
                                 device_token_hashed character varying(32) NOT NULL,
                                 account_id integer NOT NULL,
                                 region character varying(3),
                                 serial_number character varying(11)
);


ALTER TABLE public.userbase OWNER TO wiisoap;

--
-- Name: owned_titles owned_titles_pk; Type: CONSTRAINT; Schema: public; Owner: wiisoap
--

ALTER TABLE ONLY public.owned_titles
    ADD CONSTRAINT owned_titles_pk PRIMARY KEY (account_id);


--
-- Name: tickets tickets_pk; Type: CONSTRAINT; Schema: public; Owner: wiisoap
--

ALTER TABLE ONLY public.tickets
    ADD CONSTRAINT tickets_pk PRIMARY KEY (title_id);


--
-- Name: userbase userbase_pk; Type: CONSTRAINT; Schema: public; Owner: wiisoap
--

ALTER TABLE ONLY public.userbase
    ADD CONSTRAINT userbase_pk PRIMARY KEY (account_id);


--
-- Name: owned_titles_account_id_uindex; Type: INDEX; Schema: public; Owner: wiisoap
--

CREATE UNIQUE INDEX owned_titles_account_id_uindex ON public.owned_titles USING btree (account_id);


--
-- Name: userbase_account_id_uindex; Type: INDEX; Schema: public; Owner: wiisoap
--

CREATE UNIQUE INDEX userbase_account_id_uindex ON public.userbase USING btree (account_id);


--
-- Name: userbase_device_id_uindex; Type: INDEX; Schema: public; Owner: wiisoap
--

CREATE UNIQUE INDEX userbase_device_id_uindex ON public.userbase USING btree (device_id);


--
-- Name: userbase_device_token_uindex; Type: INDEX; Schema: public; Owner: wiisoap
--

CREATE UNIQUE INDEX userbase_device_token_uindex ON public.userbase USING btree (device_token);

--
-- Name: owned_titles order_account_ids; Type: FK CONSTRAINT; Schema: public; Owner: wiisoap
--

ALTER TABLE ONLY public.owned_titles
    ADD CONSTRAINT order_account_ids FOREIGN KEY (account_id) REFERENCES public.userbase(account_id);


--
-- PostgreSQL database dump complete
--