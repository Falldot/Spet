<template>
<v-layout row justify-center align-center>
    <v-card v-for="item in DOCS"
        :key="item"
        class="mx-5 scale"
        max-width="200"
        min-height="300px"
        @click="move(item)"
    >
        <v-img
        src="../assets/docx.png"
        height="200px"
        width="200px"
        ></v-img>

        <v-card-title>
        {{item}}
        </v-card-title>
    </v-card>
</v-layout>
</template>

<script>
export default {
    name: "ListDocs",
    data: () => ({
        docs: []
    }),
    computed: {
        DOCS() {
            return this.$store.state.docs.DOCS
        }
    },
    created() {
        this.$store.dispatch("GET_DOCS")
    },
    methods: {
        move(item) {
            let i = item.slice(0, -5)
            this.$router.push({ path:'/docs/edit'})
            this.$store.dispatch("GET_DOCS_NAME", i)
        }
    }
}
</script>

<style>
.spet-color {
    background: linear-gradient(to top, rgba(175, 255, 175, 0.7), #84bbff) !important;
}

.logo {
    position: absolute;
    width: 300px;
    margin: 50px;
}

.scale{
    cursor: pointer;
    transition: 0.5s;
    display: block;
}
.scale:hover {
    transform: scale(1.1);
}
</style>