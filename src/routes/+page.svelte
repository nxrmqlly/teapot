<script>
    import Logo from "$lib/components/Logo.svelte";
    import RSVP from "$lib/components/RSVP.svelte";
    import { onMount } from "svelte";
    import { gsap, SplitText } from "$lib/gsap";

    onMount(() => {
        let split = SplitText.create(".tagline-p", {
            type: "words",
        });
        gsap.from([...split.words], {
            duration: 0.5,
            y: 20, // animate from 100px below
            filter: "blur(10px)",
            autoAlpha: 0, // fade in from opacity: 0 and visibility: hidden
            stagger: 0.05, // 0.05 seconds between each
        });
    });
</script>

<section
    class="main flex items-center justify-center bg-tea-background w-screen h-screen flex-col"
>
    <div class="logo-container lg:w-1/2 md:w-2/3 w-5/6">
        <Logo />
    </div>
    <div class="tagline">
        <p
            class="tagline-p text-3xl md:text-4xl lg:text-5xl text-center font-instrument text-tea-dark-olive mx-4"
        >
            Build stuff with HTTP, <br />
            Get matcha, domains, cloud credits and more
        </p>
    </div>
    <div class="rsvp-btn mt-10">
        <a href="#a">
            <RSVP />
        </a>
    </div>
</section>
