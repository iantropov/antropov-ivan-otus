import React from 'react';
import { NextPage } from 'next';
import Link from 'next/link';

const Home: NextPage = () => {
    return (
        <section className="my-content">
            <div className="my-header">
                <h1>Hello, World 22!</h1>
            </div>
            <div className="my-body">
                <Link href="/login">Login</Link>
                <br/>
                <Link href="/register">Register</Link>
            </div>
        </section>
    );
};

export default Home;
