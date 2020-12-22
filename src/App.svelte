<script lang="ts">
    // first we check if we have received a token
    import {onMount} from "svelte";
    import {requireTwitchLogin} from "./requireTwitchLogin";

    let loading = true;
    let token = '';

    onMount(async () => {
        const user = await requireTwitchLogin();
        loading = false;
        token = user.display_name;
    })
</script>

{#if loading}
    <p>LOADING...</p>
{:else}
    <p>{token}</p>
{/if}

<style>
</style>
