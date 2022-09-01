import app from "./App.svelte"
import members from "./Members.svelte"
import groups from "./Groups.svelte"
import guests from "./Guests.svelte"
import insert from "./Insert.svelte"
import view from "./View.svelte"
import supportingdata from "./SupportingData.svelte"
import dataexport from "./DataExport.svelte"
import clinicaltrialdashboard from "./ClinicalTrialDashboard.svelte"
import newtrial from "./NewTrial.svelte"

// export other components here.
export default {
    app: app,
    members:members,
    groups: groups,
    guests: guests,
    insert: insert,
    view: view,
    supportingdata: supportingdata,
    dataexport: dataexport,
    clinicaltrialdashboard: clinicaltrialdashboard,
    newtrial: newtrial,
}
