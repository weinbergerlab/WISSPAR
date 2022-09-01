<script>export let data;</script>
<div class="tr-indent1" style="margin-top:2em; border:1px solid white">
    <div class="tr-indent2" style="margin-top:1em">

        <div style="width:98% ; margin-top:1em ">
            <div class="tr-indent2" style=""><div class="ct-header2">Results Point of Contact</div>

                <table class="ct-layout_table tr-tableStyle tr-moreInfo" style="margin-top:1ex;margin-left:1em">
                    <caption>Layout table for Results Point of Contact information</caption>
                    <thead>
                    <tr>
                        <th id="colTitle1"></th>
                        <th id="colData1"></th>
                    </tr>
                    </thead>
                    <tbody>
                    <tr>
                        <td class="colTitle1">Name/Title:</td>
                        <td headers="colData1" style="padding-left:1em">{data.ResultsSection.MoreInfoModule.PointOfContact.PointOfContactTitle}</td>
                    </tr>
                    <tr>
                        <td class="colTitle1">Organization:</td>
                        <td headers="colData1" style="padding-left:1em">{data.ResultsSection.MoreInfoModule.PointOfContact.PointOfContactOrganization}</td>
                    </tr>
                    <tr>
                        <td class="colTitle1">Phone:</td>
                        <td headers="colData1" style="padding-left:1em">{data.ResultsSection.MoreInfoModule.PointOfContact.PointOfContactPhone}</td>
                    </tr>
                    <tr>
                        <td class="colTitle1">Email:</td>
                        <td headers="colData1" style="padding-left:1em">{data.ResultsSection.MoreInfoModule.PointOfContact.PointOfContactEMail}</td>
                    </tr>
                    </tbody>
                </table>
            </div>
        </div>

        <table class="ct-layout_table tr-tableStyle tr-moreInfo" style="margin-top:1ex;margin-left:1em">
            <caption>Layout table for additonal information</caption>
            <thead>
            <tr>
                <th id="colTitle2"></th>
                <th id="colData2"></th>
            </tr>
            </thead>
            <tbody>
            <tr>
                <td class="colTitle2" >Responsible Party:</td>
                <td headers="colData2" style="padding-left:1em">{data.ProtocolSection.IdentificationModule.Organization.OrgFullName}</td>
            </tr>
            <tr>
                <td class="colTitle2">Other Study ID Numbers:</td>
                <td headers="colData2" style="padding-left:1em">{data.ProtocolSection.IdentificationModule.OrgStudyIdInfo.OrgStudyId}</td>
            </tr>
            <tr>
                <td class="colTitle2">First Submitted:</td>
                <td headers="colData2" style="padding-left:1em">{data.ProtocolSection.StatusModule.StudyFirstSubmitDate}</td>
            </tr>
            <tr>
                <td class="colTitle2">First Posted:</td>
                <td headers="colData2" style="padding-left:1em">{data.ProtocolSection.StatusModule.StudyFirstPostDateStruct.StudyFirstPostDate}</td>
            </tr>
            <tr>
                <td class="colTitle2">Results First Submitted:</td>
                <td headers="colData2" style="padding-left:1em">{data.ProtocolSection.StatusModule.ResultsFirstSubmitDate}</td>
            </tr>
            <tr>
                <td class="colTitle2">Results First Posted:</td>
                <td headers="colData2" style="padding-left:1em">{data.ProtocolSection.StatusModule.ResultsFirstPostDateStruct.ResultsFirstPostDate}</td>
            </tr>
            <tr>
                <td class="colTitle2">Last Update Posted:</td>
                <td headers="colData2" style="padding-left:1em">{data.ProtocolSection.StatusModule.LastUpdatePostDateStruct.LastUpdatePostDate}</td>
            </tr>
            </tbody>
        </table>
    </div>
</div>
