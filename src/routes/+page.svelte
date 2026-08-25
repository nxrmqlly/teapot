<script lang="ts">
    import { onMount } from "svelte";
    import Logo from "$lib/components/Logo.svelte";
    import RSVP from "$lib/components/RSVP.svelte";
    import LeafRain from "$lib/components/LeafRain.svelte";

    onMount(async () => {
        const { gsap } = await import("gsap");
        const { SplitText } = await import("gsap/SplitText");

        gsap.registerPlugin(SplitText);

        const split = SplitText.create(".tagline-p", {
            type: "words",
        });

        gsap.from(split.words, {
            duration: 0.5,
            y: 20,
            filter: "blur(10px)",
            autoAlpha: 0,
            stagger: 0.05,
        });
    });
</script>

<section
    class="main flex items-center justify-center bg-tea-background w-screen h-screen flex-col z-10"
>
    <div class="logo-container lg:w-1/2 md:w-2/3 w-5/6 z-20">
        <Logo />
    </div>

    <div class="z-20 tagline">
        <p
            class="tagline-p text-3xl md:text-4xl lg:text-5xl text-center font-instrument text-tea-dark-olive mx-4"
        >
            Build stuff with HTTP, <br />
            Get matcha, domains, cloud credits and more
        </p>
    </div>

    <div class="z-20 rsvp-btn mt-10">
        <a href="#a">
            <RSVP />
        </a>
    </div>
</section>

<LeafRain />
