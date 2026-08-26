<script lang="ts">
    import { onMount } from "svelte";
    import Check from "@lucide/svelte/icons/check";
    import X from "@lucide/svelte/icons/x";
    import Logo from "$lib/components/Logo.svelte";
    import Card from "$lib/components/Card.svelte";
    import RSVP from "$lib/components/RSVP.svelte";
    import LeafRain from "$lib/components/LeafRain.svelte";

    const make = [
        "rest api",
        "reverse proxy",
        "load balancer",
        "http tunnels",
        "api wrapper/sdk",
        "custom protocol based on http",
        "...or anything else!",
    ];
    const dont = ["a static html website", "a frontend calling an api"];

    onMount(async () => {
        const { gsap } = await import("gsap");
        const { SplitText } = await import("gsap/SplitText");
        const { ScrollTrigger } = await import("gsap/ScrollTrigger");

        gsap.registerPlugin(SplitText, ScrollTrigger);

        const split = SplitText.create(".tagline-p", {
            type: "words",
        });

        const blurin = {
            duration: 0.5,
            y: 40,
            filter: "blur(10px)",
            autoAlpha: 0,
            stagger: 0.05,
        };

        gsap.from(split.words, blurin);

        const tl = gsap.timeline({
            scrollTrigger: {
                trigger: ".info-sec",
                start: "top 85%",
                end: "top 20%",
                scrub: 1,
            },
        });

        const statSplit = SplitText.create(".teapot-stat-heading", {
            type: "words",
        });

        tl.from(statSplit.words, blurin);
        tl.from(".proj-req-card", blurin);
    });
</script>

<section
    class="main relative top-0 flex items-center justify-center bg-tea-background z-10 w-screen h-screen flex-col"
>
    <LeafRain />

    <div
        class="logo-container lg:w-1/2 md:w-2/3 w-5/6 z-10 hover:scale-105 transition-transform"
    >
        <Logo />
    </div>

    <div class="z-10 tagline">
        <p
            class="tagline-p text-3xl md:text-4xl lg:text-5xl text-center font-instrument text-tea-dark-olive mx-4"
        >
            Build stuff with HTTP, <br />
            Get matcha, domains, cloud credits and more
        </p>
    </div>

    <div class="z-10 rsvp-btn mt-10">
        <a
            href="https://rsvp.soon.it/tea/"
            rel="noopener noreferrer"
            target="_blank"
        >
            <RSVP />
        </a>
    </div>
</section>

<section class="content-set bg-tea-pale-green">
    <h1
        class="italic teapot-stat-heading font-instrument text-7xl text-center pt-20 text-tea-dark-olive"
    >
        What can I build?
    </h1>
    <div class="project-spec-container grid w-160 h-90 gap-3 my-20">
        <div class="proj-req-card make w-full font-inst-sans italic">
            <Card title="make something that speaks http">
                {#each make as m}
                    <div class="text-2xl pl-9 -indent-9 leading-7">
                        <Check
                            size="28"
                            class="inline-block align-text-bottom text-tea-olive"
                        />
                        {m}
                    </div>
                {/each}
            </Card>
        </div>
        <div class="proj-req-card dont w-full font-inst-sans italic">
            <Card title="but don't make">
                {#each dont as m}
                    <div class="text-2xl pl-9 -indent-9 leading-7">
                        <X
                            size="28"
                            class="inline-block align-text-bottom text-tea-berry"
                        />
                        {m}
                    </div>
                {/each}
            </Card>
        </div>
        <div class="proj-req-card bonus w-full font-inst-sans italic">
            <Card title="bonus!">
                <p class="text-2xl leading-6">
                    implement the “418 I’m a teapot” status in your project and
                    deploy it for extra stuff!
                </p>
            </Card>
        </div>
    </div>
</section>

<section class="content-set bg-tea-background">
    <div class="font-instrument text-3xl italic">hi this is wip</div>
</section>

<section class="content-set bg-tea-berry">
    <div class="font-instrument text-3xl italic">hi this is wip</div>
</section>

<footer
    class="content-set h-60! bg-tea-olive font-instrument text-3xl italic text-tea-pale-green"
>
    insert footer here
</footer>

<style>
    .project-spec-container {
        grid-template-areas:
            "make dont"
            "make bonus";
        grid-template-columns: 1fr 1fr;
    }
    .make {
        grid-area: make;
    }
    .dont {
        grid-area: dont;
    }
    .bonus {
        grid-area: bonus;
    }

    .content-set {
        position: relative;
        display: flex;
        flex-direction: column;
        align-items: center;
        width: 100vw;
        height: 100vh;
        position: sticky;
        top: 0;
        z-index: 20;
    }

    .content-set::before {
        content: "";
        position: absolute;
        top: -3rem;
        left: 0;
        width: 100%;
        height: 3rem;
        border-radius: 3rem 3rem 0 0;
        background: inherit;
    }
</style>
