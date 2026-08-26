<script lang="ts">
    import "./layout.css";
    import favicon from "$lib/assets/favicon.svg";
    import Lenis from "lenis";
    import { ScrollTrigger } from "gsap/ScrollTrigger";
    import { gsap } from "gsap";
    import { onMount } from "svelte";

    onMount(() => {
        const lenis = new Lenis({
            autoRaf: true,
        });
        lenis.on("scroll", ScrollTrigger.update);
        gsap.ticker.add((time) => {
            lenis.raf(time * 1000);
        });
        gsap.ticker.lagSmoothing(0);
    });

    let { children } = $props();
</script>

<svelte:head>
    <link rel="icon" href={favicon} />
</svelte:head>
{@render children()}
