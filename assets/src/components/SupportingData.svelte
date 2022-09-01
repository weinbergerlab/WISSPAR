<script>
    import {slide} from "svelte/transition";
    import {elasticInOut} from "svelte/easing";
    import {onMount} from 'svelte';

    let vaccines = [];
    let manufacturers = [];
    let dose_descriptions = [];
    let schedules = [];
    let immunocompromised_groups = [];
    let use_cases = [];
    let vaccineName = "";
    let manufacturerName = "";
    let doseDescriptionName = "";
    let scheduleName = "";
    let immunocompromisedGroupName = "";
    let useCaseName = "";
    let useCaseDescription = "";
    let useCaseCode = "";

    const vaccinesAPI = '/api/vaccine';
    const manufacturersAPI = '/api/manufacturer';
    const doseDescriptionAPI = '/api/dose-description';
    const scheduleAPI = '/api/schedule';
    const immunocompromisedGroupsAPI = '/api/icg';
    const useCaseAPI = '/api/usecase';

    onMount(async () => {
        const res = await fetch(vaccinesAPI);
        vaccines = await res.json();

        const resManufacturer = await fetch(manufacturersAPI);
        manufacturers = await resManufacturer.json();

        const resDoseDescription = await fetch(doseDescriptionAPI);
        dose_descriptions = await resDoseDescription.json();

        const resSchedule = await fetch(scheduleAPI);
        schedules = await resSchedule.json();

        const resICG = await fetch(immunocompromisedGroupsAPI);
        immunocompromised_groups = await resICG.json();

        const resUseCase = await fetch(useCaseAPI);
        use_cases = await resUseCase.json();
    });

    const addVaccine = async() =>{
        if (!vaccineName){
            return
        }

        const res = await fetch(vaccinesAPI, {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify({
                name: vaccineName
            })
        });
        const vaccine = await res.json();
        const resVaccines = await fetch(vaccinesAPI);
        vaccines = await resVaccines.json();
        vaccineName = "";
    }

    const updateVaccine = async(id, name) => {
        const res = await fetch(vaccinesAPI +"/" + id, {
            method: 'PUT',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify({
                name: name
            })
        });
        const vaccine = await res.json();
        const resVaccines = await fetch(vaccinesAPI);
        vaccines = await resVaccines.json();
    }

    const removeVaccine = async(id) => {
        const res = await fetch(vaccinesAPI +"/" + id, {
            method: 'DELETE',
            headers: {'Content-Type': 'application/json'}
        });
        const resVaccines = await fetch(vaccinesAPI);
        vaccines = await resVaccines.json();
    }

    const addManufacturer = async() =>{
        if (!manufacturerName){
            return
        }

        const res = await fetch(manufacturersAPI, {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify({
                name: manufacturerName
            })
        });
        const manufacturer = await res.json();
        const resManufacturer = await fetch(manufacturersAPI);
        manufacturers = await resManufacturer.json();
        manufacturerName = "";
    }

    const updateManufacturer = async(id, name) => {
        const res = await fetch(manufacturersAPI +"/" + id, {
            method: 'PUT',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify({
                name: name
            })
        });
        const manufacturer = await res.json();
        const resManufacturer = await fetch(manufacturersAPI);
        manufacturers = await resManufacturer.json();
    }

    const removeManufacturer = async(id) => {
        const res = await fetch(manufacturersAPI +"/" + id, {
            method: 'DELETE',
            headers: {'Content-Type': 'application/json'}
        });
        const manufacturer = await res.json();
        const resManufacturer = await fetch(manufacturersAPI);
        manufacturers = await resManufacturer.json();
    }

    const addDoseDescription = async() =>{
        if (!doseDescriptionName){
            return
        }

        const res = await fetch(doseDescriptionAPI, {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify({
                name: doseDescriptionName
            })
        });
        const dose_description = await res.json();
        const resDoseDescription = await fetch(doseDescriptionAPI);
        dose_descriptions = await resDoseDescription.json();
        doseDescriptionName = "";
    }

    const updateDoseDescription = async(id, name) => {
        const res = await fetch(doseDescriptionAPI +"/" + id, {
            method: 'PUT',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify({
                name: name
            })
        });
        const dose_description = await res.json();
        const resDoseDescription = await fetch(doseDescriptionAPI);
        dose_descriptions = await resDoseDescription.json();
    }

    const removeDoseDescription = async(id) => {
        const res = await fetch(doseDescriptionAPI +"/" + id, {
            method: 'DELETE',
            headers: {'Content-Type': 'application/json'}
        });
        const dose_description = await res.json();
        const resDoseDescription = await fetch(doseDescriptionAPI);
        dose_descriptions = await resDoseDescription.json();
    }

    const addSchedule = async() =>{
        if (!scheduleName){
            return
        }

        const res = await fetch(scheduleAPI, {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify({
                name: scheduleName
            })
        });
        const schedule = await res.json();
        const resSchedule = await fetch(scheduleAPI);
        schedules = await resSchedule.json();
        scheduleName = "";
    }

    const updateSchedule = async(id, name) => {
        const res = await fetch(scheduleAPI +"/" + id, {
            method: 'PUT',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify({
                name: name
            })
        });
        const schedule = await res.json();
        const resSchedule = await fetch(scheduleAPI);
        schedules = await resSchedule.json();
    }

    const removeSchedule = async(id) => {
        const res = await fetch(scheduleAPI +"/" + id, {
            method: 'DELETE',
            headers: {'Content-Type': 'application/json'}
        });
        const schedule = await res.json();
        const resSchedule = await fetch(scheduleAPI);
        schedules = await resSchedule.json();
    }

    const addICG = async() =>{
        if (!immunocompromisedGroupName){
            return
        }

        const res = await fetch(immunocompromisedGroupsAPI, {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify({
                name: immunocompromisedGroupName
            })
        });
        const icg = await res.json();
        const resICG = await fetch(immunocompromisedGroupsAPI);
        immunocompromised_groups = await resICG.json();
        immunocompromisedGroupName = "";
    }

    const updateICG = async(id, name) => {
        const res = await fetch(immunocompromisedGroupsAPI +"/" + id, {
            method: 'PUT',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify({
                name: name
            })
        });
        const icg = await res.json();
        const resICG = await fetch(immunocompromisedGroupsAPI);
        immunocompromised_groups = await resICG.json();
    }

    const removeICG = async(id) => {
        const res = await fetch(immunocompromisedGroupsAPI +"/" + id, {
            method: 'DELETE',
            headers: {'Content-Type': 'application/json'}
        });
        const icg = await res.json();
        const resICG = await fetch(immunocompromisedGroupsAPI);
        immunocompromised_groups = await resICG.json();
    }

    const addUseCase = async() =>{
        if (!useCaseName || !useCaseCode){
            return
        }

        const res = await fetch(useCaseAPI, {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify({
                name: useCaseName,
                description: useCaseDescription,
                code: useCaseCode
            })
        });
        const useCase = await res.json();
        const resUseCase = await fetch(useCaseAPI);
        use_cases = await resUseCase.json();
        useCaseName = "";
        useCaseDescription = "";
        useCaseCode = "";
    }

    const updateUseCase = async(id, name, description, code) => {
        const res = await fetch(useCaseAPI +"/" + id, {
            method: 'PUT',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify({
                name: name,
                description: description,
                code: code
            })
        });
        const useCase = await res.json();
        const resUseCases = await fetch(useCaseAPI);
        use_cases = await resUseCases.json();
    }

    const removeUseCase = async(id) => {
        const res = await fetch(useCaseAPI +"/" + id, {
            method: 'DELETE',
            headers: {'Content-Type': 'application/json'}
        });
        const resUseCases = await fetch(useCaseAPI);
        use_cases = await resUseCases.json();
    }
</script>

<main class="container is-fluid">
    <h1 style="text-align: center" class="title">Admin Section</h1>
    <div class="columns is-centered is-mobile" >
        <div class="column" style="width: 70%">
            <h3 class="has-text-centered subtitle">Vaccines</h3>
            <form class="field has-addons" style="justify-content: center" on:submit|preventDefault={addVaccine}>
                <div class="control">
                    <input bind:value={vaccineName} class="input" type="text" placeholder="Vaccine Name">
                </div>
                <div class="control">
                    <button class="button is-primary">
                        <span class="icon is-small">
                          <i class="fas fa-plus"></i>
                        </span>
                    </button>
                </div>
            </form>
            <ul class:list={vaccines.length > 0}>
                {#each vaccines as vaccine (vaccine.ID)}
                    <li class="list-item" transition:slide="{{duration: 300, easing: elasticInOut}}">
                        <div class="is-flex" style="align-items: center">
                            <input bind:value={vaccine.Name} class="input" type="text">
                            <div style="flex: 1"></div>
                            <button class="button is-text is-pulled-right is-small" on:click={()=> updateVaccine(vaccine.ID, vaccine.Name)}>
                                <span class="icon">
                                  <i class="fas fa-save"></i>
                                </span>
                            </button>
                            <button class="button is-text is-pulled-right is-small" on:click={()=> removeVaccine(vaccine.ID)}>
                                <span class="icon">
                                  <i class="fas fa-trash"></i>
                                </span>
                            </button>
                        </div>
                    </li>
                {:else}
                    <li class="has-text-centered"
                        transition:slide="{{delay: 600, duration: 300, easing: elasticInOut}}">No vaccines yet, please add them.
                    </li>
                {/each}
            </ul>
        </div>
        <div class="column" style="width: 70%">
            <h3 class="has-text-centered subtitle">Manufacturers</h3>
            <div class="field has-addons" style="justify-content: center">
                <div class="control">
                    <input bind:value={manufacturerName} class="input" type="text" placeholder="Manufacturer Name">
                </div>
                <div class="control">
                    <button type="button" class="button is-primary" on:click={() => addManufacturer()}>
                        <span class="icon is-small">
                            <i class="fas fa-plus"></i>
                        </span>
                    </button>
                </div>
            </div>
            <ul class:list={manufacturers.length > 0}>
                {#each manufacturers as manufacturer (manufacturer.ID)}
                    <li class="list-item" transition:slide="{{duration: 300, easing: elasticInOut}}">
                        <div class="is-flex" style="align-items: center">
                            <input bind:value={manufacturer.Name} class="input" type="text">
                            <div style="flex: 1"></div>
                            <button class="button is-text is-pulled-right is-small" on:click={()=> updateManufacturer(manufacturer.ID, manufacturer.Name)}>
                                <span class="icon"><i class="fas fa-save"></i></span>
                            </button>
                            <button class="button is-text is-pulled-right is-small" on:click={()=> removeManufacturer(manufacturer.ID)}>
                                <span class="icon"><i class="fas fa-trash"></i></span>
                            </button>
                        </div>
                    </li>
                {:else}
                    <li class="has-text-centered"
                        transition:slide="{{delay: 600, duration: 300, easing: elasticInOut}}">No manufacturers yet, please add them.
                    </li>
                {/each}
            </ul>
        </div>
        <div class="column" style="width: 70%">
            <h3 class="has-text-centered subtitle">Immunocompromised Groups</h3>
            <div class="field has-addons" style="justify-content: center">
                <div class="control">
                    <input bind:value={immunocompromisedGroupName} class="input" type="text" placeholder="Manufacturer Name">
                </div>
                <div class="control">
                    <button type="button" class="button is-primary" on:click={() => addICG()}>
                        <span class="icon is-small">
                            <i class="fas fa-plus"></i>
                        </span>
                    </button>
                </div>
            </div>
            <ul class:list={immunocompromised_groups.length > 0}>
                {#each immunocompromised_groups as icg (icg.ID)}
                    <li class="list-item" transition:slide="{{duration: 300, easing: elasticInOut}}">
                        <div class="is-flex" style="align-items: center">
                            <input bind:value={icg.Name} class="input" type="text">
                            <div style="flex: 1"></div>
                            <button class="button is-text is-pulled-right is-small" on:click={()=> updateICG(icg.ID, icg.Name)}>
                                <span class="icon"><i class="fas fa-save"></i></span>
                            </button>
                            <button class="button is-text is-pulled-right is-small" on:click={()=> removeICG(icg.ID)}>
                                <span class="icon"><i class="fas fa-trash"></i></span>
                            </button>
                        </div>
                    </li>
                {:else}
                    <li class="has-text-centered"
                        transition:slide="{{delay: 600, duration: 300, easing: elasticInOut}}">No immunocompromised groups yet, please add them.
                    </li>
                {/each}
            </ul>
        </div>
    </div>
    <hr />
    <div class="columns">
    <div class="column" style="width: 70%">
        <h3 class="has-text-centered subtitle">Dose Descriptions</h3>
        <div class="field has-addons" style="justify-content: center">
            <div class="control">
                <input bind:value={doseDescriptionName} class="input" type="text" placeholder="Dose Description Name">
            </div>
            <div class="control">
                <button type="button" class="button is-primary" on:click={() => addDoseDescription()}>
                        <span class="icon is-small">
                            <i class="fas fa-plus"></i>
                        </span>
                </button>
            </div>
        </div>
        <ul class:list={dose_descriptions.length > 0}>
            {#each dose_descriptions as dose_description (dose_description.ID)}
                <li class="list-item" transition:slide="{{duration: 300, easing: elasticInOut}}">
                    <div class="is-flex" style="align-items: center">
                        <input bind:value={dose_description.Name} class="input" type="text">
                        <div style="flex: 1"></div>
                        <button class="button is-text is-pulled-right is-small" on:click={()=> updateDoseDescription(dose_description.ID, dose_description.Name)}>
                            <span class="icon"><i class="fas fa-save"></i></span>
                        </button>
                        <button class="button is-text is-pulled-right is-small" on:click={()=> removeDoseDescription(dose_description.ID)}>
                            <span class="icon"><i class="fas fa-trash"></i></span>
                        </button>
                    </div>
                </li>
            {:else}
                <li class="has-text-centered"
                    transition:slide="{{delay: 600, duration: 300, easing: elasticInOut}}">No dose descriptions yet, please add them.
                </li>
            {/each}
        </ul>
    </div>
    <div class="column" style="width: 70%">
        <h3 class="has-text-centered subtitle">Schedules</h3>
        <div class="field has-addons" style="justify-content: center">
            <div class="control">
                <input bind:value={scheduleName} class="input" type="text" placeholder="Schedule Name">
            </div>
            <div class="control">
                <button type="button" class="button is-primary" on:click={() => addSchedule()}>
                        <span class="icon is-small">
                            <i class="fas fa-plus"></i>
                        </span>
                </button>
            </div>
        </div>
        <ul class:list={schedules.length > 0}>
            {#each schedules as schedule (schedule.ID)}
                <li class="list-item" transition:slide="{{duration: 300, easing: elasticInOut}}">
                    <div class="is-flex" style="align-items: center">
                        <input bind:value={schedule.Name} class="input" type="text">
                        <div style="flex: 1"></div>
                        <button class="button is-text is-pulled-right is-small" on:click={()=> updateSchedule(schedule.ID, schedule.Name)}>
                            <span class="icon"><i class="fas fa-save"></i></span>
                        </button>
                        <button class="button is-text is-pulled-right is-small" on:click={()=> removeSchedule(schedule.ID)}>
                            <span class="icon"><i class="fas fa-trash"></i></span>
                        </button>
                    </div>
                </li>
            {:else}
                <li class="has-text-centered"
                    transition:slide="{{delay: 600, duration: 300, easing: elasticInOut}}">No schedules yet, please add them.
                </li>
            {/each}
        </ul>
    </div>
</div>
    <hr />
    <div class="columns">
        <div class="column" style="width: 70%">
            <h2 class="has-text-centered title">Use Cases</h2>
            <div class="is-flex" style="align-items: center">
                <input bind:value={useCaseName} class="input" type="text" placeholder="Use Case Name">
                <input bind:value={useCaseDescription} class="input" type="text" placeholder="Use Case Description">
                <input bind:value={useCaseCode} class="input" type="text" placeholder="Use Case Code (no spaces)">
                <button type="button" class="button is-primary is-pulled-right is-small" on:click={() => addUseCase()}>
                    <span class="icon">
                        <i class="fas fa-plus"></i>
                    </span>
                </button>
                <button type="button" class="button is-pulled-right is-small" style="visibility: hidden">
                    <span class="icon">
                    </span>
                </button>
            </div>
            <ul class:list={use_cases.length > 0}>
                {#each use_cases as use_case (use_case.ID)}
                    <li class="list-item" transition:slide="{{duration: 300, easing: elasticInOut}}">
                        <div class="is-flex" style="align-items: center">
                            <input bind:value={use_case.Name} class="input" type="text">
                            <input bind:value={use_case.Description} class="input" type="text">
                            <input bind:value={use_case.Code} class="input" type="text">
                            <div style="flex: 1"></div>
                            <button class="button is-text is-pulled-right is-small" on:click={()=> updateUseCase(use_case.ID, use_case.Name, use_case.Description, use_case.Code)}>
                                <span class="icon"><i class="fas fa-save"></i></span>
                            </button>
                            <button class="button is-text is-pulled-right is-small" on:click={()=> removeUseCase(use_case.ID)}>
                                <span class="icon"><i class="fas fa-trash"></i></span>
                            </button>
                        </div>
                    </li>
                {:else}
                    <li class="has-text-centered"
                        transition:slide="{{delay: 600, duration: 300, easing: elasticInOut}}">No use cases yet, please add them.
                    </li>
                {/each}
            </ul>
        </div>
    </div>
</main>
