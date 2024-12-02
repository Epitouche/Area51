<template>
    <Navbar/>
    <main>
        <h1>Info of n°{{ post.id }}</h1>
        <h2>{{ post.title }}</h2>
        <p>
            {{ post.body }}
        </p>
    </main>
</template>

<script setup>
const route = useRoute()
const { data : post } = useFetch(() => "https://jsonplaceholder.typicode.com/posts/" + route.params.id)
</script>
