<script lang="ts">
    const steps = [
        {
            num: "01",
            title: "join the community",
            lead: "Hop into our Slack channel.",
            sub: "come say hi, lurk quietly, or accidentally make friends. up to you!",
        },
        {
            num: "02",
            title: "build something",
            lead: "Website, game, tool, app, whatever.",
            sub: "the only rule? it has to run over HTTP.",
        },
        {
            num: "03",
            title: "track your time",
            lead: "Use Hackatime while you build.",
            sub: "yes, we need proof you actually made the thing and didn't just stare at VS Code for 6 hours.",
        },
        {
            num: "04",
            title: "send it in",
            lead: "Submit it once it runs.",
            sub: "show us what you built, how long it took, and hopefully something that actually works.",
        },
        {
            num: "05",
            title: "get your prize",
            lead: "Matcha, domains, cloud credits and more.",
            sub: "sleep schedule not included.",
        },
    ];
</script>

<section
    id="how-it-works"
    class="w-full bg-tea-minty-green px-4 py-20 md:py-28"
>
    <div class="mx-auto max-w-5xl">
        <h2 class="reveal font-instrument text-4xl text-tea-dark-olive md:text-6xl">
            ✦ how it works
        </h2>

        <div class="mt-12 grid gap-x-10 gap-y-14 md:mt-16 md:grid-cols-2 lg:grid-cols-3">
            {#each steps as step (step.num)}
                <div class="reveal">
                    <p class="font-instrument text-6xl text-tea-olive md:text-7xl">
                        {step.num}
                    </p>
                    <h3 class="font-instrument mt-3 text-2xl text-tea-dark-olive md:text-3xl">
                        {step.title}
                    </h3>
                    <p class="mt-2 text-lg text-tea-dark-olive">{step.lead}</p>
                    <p class="mt-1 text-base text-tea-olive italic">{step.sub}</p>
                </div>
            {/each}
        </div>
    </div>
</section>
